package templates

type templateBase struct {
	Params  []string
	Content string
}

type templates map[string]templateBase

var templateTypes = templates{
	"test": templateBase{
		Params:  []string{"phone"},
		Content: "🤖 Teste de envio para o seu numero {{.phone}}",
	},

	"greetings": templateBase{
		Params:  []string{"phone", "name"},
		Content: "Oi {{.name}}, tudo bem?\nVocê foi adicionado como contato na minha lista de notificações 🤘. Se eu encontrar algum ingresso para você, pode deixar que eu aviso! 😎",
	},

	"found_tickets": templateBase{
		Params:  []string{"phone", "name", "eventName", "url"},
		Content: "Ei {{.name}}, encontrei ingressos para o evento {{.eventName}} 😱!!! Fica esperto, aqui o link: {{.url}}",
	},
}
