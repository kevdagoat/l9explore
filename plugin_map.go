package l9explore

import (
	"github.com/LeakIX/l9format"
	// Import your plugins here
	"github.com/LeakIX/l9plugins"
	//l9_nuclei_plugin "github.com/gboddin/l9-nuclei-plugin"
	metabase_plugin "github.com/kaizensecurity/CVE-2021-41277"
	kevdagoats_plugins "github.com/kevdagoat/leakix-plugins"
)

var TcpPlugins []l9format.ServicePluginInterface
var WebPlugins []l9format.WebPluginInterface

func LoadL9ExplorePlugins() {
	TcpPlugins = append(TcpPlugins, l9plugins.GetTcpPlugins()...)
	WebPlugins = append(WebPlugins, l9plugins.GetWebPlugins()...)
	// Add your plugins here
	//TcpPlugins = append(TcpPlugins, l9_nuclei_plugin.NucleiPlugin{})
	WebPlugins = append(WebPlugins, metabase_plugin.MetabaseHttpPlugin{})
	TcpPlugins = append(TcpPlugins, kevdagoats_plugins.FTPOpenPlugin{})
	TcpPlugins = append(TcpPlugins, kevdagoats_plugins.MemcachedOpenPlugin{})
}

