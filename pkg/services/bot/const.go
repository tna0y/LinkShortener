package bot

// Could have done this with templates

const (
	serviceName = "bot"

	helpMessage = `Welcome to the link shortener!
Commands:
/add <LINK ID> <LINK TARGET> [TIMEOUT IN SECONDS]
/list
/remove <LINK ID>`

	linkCreatedMessage = "Link created: %s"

	linkListMessage = "You have %d links:\n%s"

	linkListItemMessage = "%s -> %s %s\n"

	linkListExpiresMessage = "(expires at %s)"

	linkRemovedMessage = "Link removed"
)
