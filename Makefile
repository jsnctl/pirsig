build:
	go test ./.../test
	go build
	./pirsig

	ffplay -f f32le -ar 44100 -showmode 1 out.bin


test:
	go test ./.../test
