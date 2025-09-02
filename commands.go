package commands

func init() {

	New("timezone", "get").
		Returns(
			StringValue("timezone"),
		)
	New("timezone", "set").
		Accepts(
			StringArg("timezone").
				WithDocsDefault("Europe/Moscow").
				WithDescription("Timezone to set to in ISO 8601 format"),
		).Returns(
		StringValue("timezone"),
	)

	New("ntp", "get-list").NotImplemented()
	New("ntp", "add").NotImplemented().
		Accepts(StringArg("server")).NotImplemented()
	New("ntp", "remove").
		Accepts(StringArg("server"))

	New("firewall", "get-status")
	New("firewall", "enable")
	New("firewall", "disable")

	New("modems", "get-list")
	New("modems", "enable").
		Accepts(StringArg("modem"))
	New("modems", "disable").
		Accepts(StringArg("modem"))
	New("modems", "get-signal").
		Accepts(StringArg("modem"))

	New("system", "reboot")

	New("ssh", "get-status")
	New("ssh", "enable")
	New("ssh", "disable")

	New("journals", "get-journal").
		Accepts(
			StringArg("journal").
				WithValues("system", "core", "connection", "port-forwarding"),
			IntArg("lines").
				WithDocsDefault(100),
		)

}
