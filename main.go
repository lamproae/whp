package main

import (
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/plugins/wxweb/cleaner"
	"github.com/songtianyi/wechat-go/plugins/wxweb/faceplusplus"
	"github.com/songtianyi/wechat-go/plugins/wxweb/gifer"
	"github.com/songtianyi/wechat-go/plugins/wxweb/joker"
	"github.com/songtianyi/wechat-go/plugins/wxweb/laosj"
	"github.com/songtianyi/wechat-go/plugins/wxweb/meinv"
	"github.com/songtianyi/wechat-go/plugins/wxweb/replier"
	"github.com/songtianyi/wechat-go/plugins/wxweb/revoker"
	"github.com/songtianyi/wechat-go/plugins/wxweb/switcher"
	"github.com/songtianyi/wechat-go/wxweb"
)

func main() {
	// create session
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}

	// add plugins for this session, they are disabled by default
	faceplusplus.Register(session)
	replier.Register(session)
	switcher.Register(session)
	gifer.Register(session)
	laosj.Register(session)
	cleaner.Register(session)
	revoker.Register(session)
	joker.Register(session)
	meinv.Register(session)

	// enable plugin
	session.HandlerRegister.EnableByName("switcher")
	session.HandlerRegister.EnableByName("joker")
	session.HandlerRegister.EnableByName("revoker")
	session.HandlerRegister.EnableByName("replier")
	session.HandlerRegister.EnableByName("laosj")
	session.HandlerRegister.EnableByName("gifer")
	session.HandlerRegister.EnableByName("cleaner")
	session.HandlerRegister.EnableByName("faceplusplus")
	session.HandlerRegister.EnableByName("meinv")

	if err := session.LoginAndServe(false); err != nil {
		logs.Error("session exit, %s", err)
	}
}
