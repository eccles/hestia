for p in golang kubectl skaffold helm golangci-lint mockery nodejs bats shellcheck shfmt buf just protoc
do
	asdf plugin add $p
	asdf install $p latest
done
