package templates

type templateBase struct {
	Params  []string
	Content string
}

type templates map[string]templateBase

var templateTypes = templates{
	"test": templateBase{
		Params:  []string{"phone"},
		Content: "Teste de envio parao seu numero {{.phone}}",
	},

	"greetings": templateBase{
		Params: []string{"phone", "name"},
		Content: "Olá {{.name}} tudo bem?\nVocê foi adicionado como contato" +
			" na minha lista de notificações. Se eu encontrar algum ingresso pra você pode deixar que eu aviso!",
	},

	"found_tickets": templateBase{
		Params: []string{"phone", "name", "eventName", "url"},
		Content: "Eeei {{.name}}, encontrei ingressos para o evento {{.eventName}}!!!" +
			"fica esperto, aqui o link: {{.url}}",
	},
}
