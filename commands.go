package commands

var (
	CommandTimezoneGet = New("timezone", "get").
				WithDescription("Returns device's current timezone.").
				Returns(
			StringValue("timezone"),
		)
	CommandTimezoneSet = New("timezone", "set").
				WithDescription("Sets device current timezone and returns new value.").
				Accepts(
			StringArg("timezone").
				WithDocsDefault("Europe/Moscow").
				WithDescription("Timezone to set to in ISO 8601 format"),
		).
		Returns(
			StringValue("timezone"),
		)
)

func init() {

	New("ntp", "get-list").NotReady()
	New("ntp", "add").NotReady().
		Accepts(StringArg("server")).
		NotReady()
	New("ntp", "remove").
		Accepts(StringArg("server")).
		NotReady()

	New("firewall", "get-status").NotReady()
	New("firewall", "enable").NotReady()
	New("firewall", "disable").NotReady()

	New("modems", "get-list").NotReady()
	New("modems", "enable").
		Accepts(StringArg("modem")).
		NotReady()
	New("modems", "disable").
		Accepts(StringArg("modem")).
		NotReady()
	New("modems", "get-signal").
		Accepts(StringArg("modem")).
		NotReady()

	New("system", "reboot").NotReady()

	New("ssh", "get-status").NotReady()
	New("ssh", "enable").NotReady()
	New("ssh", "disable").NotReady()

	New("journals", "get-journal").
		Accepts(
			StringArg("journal").
				WithValues("system", "core", "connection", "port-forwarding"),
			IntArg("lines").
				WithDocsDefault(100),
		).
		NotReady()

}
