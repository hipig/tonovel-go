package middlewares

import (
	"tonovel/bootstrap"
)

func Configure(b *bootstrap.Bootstrapper)  {
	cors := cors.New()
	identity := identity.New(b)
	b.UseGlobal(cors, identity)
}
