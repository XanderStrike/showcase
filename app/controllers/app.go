package controllers

import (
    "github.com/revel/revel"
    "fmt"
    "io/ioutil"
    "regexp"
    "os"
    "net/http"
    "net/url"
)

const (
    _      = iota
    KB int = 1 << (10 * iota)
    MB
    GB
)

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
    return c.Render()
}

func (c App) Login() revel.Result {
    delete(c.Session, "user")
    return c.Render()
}

func (c App) DoLogin(username, password string) revel.Result {
    user := os.Getenv("USERNAME")
    pass := os.Getenv("PASSWORD")

    if (username == user && password == pass) {
        c.Session["user"] = user
        c.Flash.Success("Successfully logged in.")
        return c.Redirect((*App).Index)
    }

    c.Flash.Error("Incorrect username or password.")
    return c.Redirect((*App).Login)
}

func (c App) CheckLogin() revel.Result {
    user := os.Getenv("USERNAME")

    if len(user) == 0 {
        return nil
    }

    if c.Action == "App.Login" || c.Action == "App.DoLogin" {
        return nil
    }

    if c.Session["user"] == user {
        return nil
    }

    c.Flash.Error("You need to log in.")
    return c.Redirect((*App).Login)
}

func (c App) Upload(torrent []byte, folder string) revel.Result {
    c.Validation.Required(torrent)
    c.Validation.MaxSize(torrent, 2*MB).
        Message("File cannot be larger than 2MB")

    c.Validation.Required(folder)
    c.Validation.Match(folder, regexp.MustCompile("(movies|music|tv|misc)"))

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect((*App).Index)
    }

    filename := c.Params.Files["torrent"][0].Filename

    err := ioutil.WriteFile(fmt.Sprintf("/watch/%s/%s", folder, filename), torrent, 0644)
    if err != nil { panic(err) }

    discord_url := os.Getenv("DISCORD_URL")

    if len(discord_url) > 0 {
    	message := fmt.Sprintf("Uploaded %s to %s folder.", filename, folder)
        resp, err := http.PostForm(discord_url, url.Values{"content":{message}, "username":{"Manual Upload"}, "avatar_url":{"https://xanderstrike.com/manual.png"}})
        if err != nil { panic(err) }
        fmt.Println(resp)
    }

    c.Flash.Success(fmt.Sprintf("Successfully uploaded to %s!", folder))
    return c.Redirect((*App).Index)
}
