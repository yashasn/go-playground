module Example/hello-world

go 1.21.1

replace Example/greetings => ../greetings //since it's a local module and it's not published

require (
	Example/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)


