package main

import (
	"github.com/mattermost/mattermost-server/v6/plugin"
	"github.com/mattermost/mattermost-server/v6/model"
	//"github.com/pkg/errors"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) UserHasBeenCreated(c *plugin.Context, user *model.User){
	config := p.API.GetConfig()
	notifyDesktop := interface{}(config.PluginSettings.Plugins["org.ishidatami.mattermost.default-settings-for-user-notify"]["defaultnotifyondesktop"]).(string)
	notifyPush := interface{}(config.PluginSettings.Plugins["org.ishidatami.mattermost.default-settings-for-user-notify"]["defaultnotifyonpush"]).(string)
	user.NotifyProps[model.DesktopNotifyProp] = notifyDesktop
	user.NotifyProps[model.PushNotifyProp] = notifyPush
	p.API.UpdateUser(user)
}

func main(){
	plugin.ClientMain(&Plugin{})
}