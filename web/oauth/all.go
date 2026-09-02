package oauth

import (
	_ "github.com/raymao96/komari/web/oauth/factory"
	_ "github.com/raymao96/komari/web/oauth/generic"
	_ "github.com/raymao96/komari/web/oauth/github"
	_ "github.com/raymao96/komari/web/oauth/qq"
)

func All() {
	//empty function to ensure all OIDC providers are registered
}
