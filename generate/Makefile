OUTPUT_DIR ?= "../fsdata"

obj = $(shell seq 10 10 250)

all: build

build: $(addprefix build/,$(obj))

$(addprefix build/,$(obj)):
	$(eval WORKDIR := $(shell mktemp -d))
	mkdir -p "$(WORKDIR)" "$(OUTPUT_DIR)"
	truncate "$(WORKDIR)/$(subst build/,,$@)g" --size "$(subst build/,,$@)g"
	mkfs.ext4 -F "$(WORKDIR)/$(subst build/,,$@)g"
	go run . --input="$(WORKDIR)/$(subst build/,,$@)g" --size="$(subst build/,,$@)" --output-dir="$(OUTPUT_DIR)"
	rm -rf $(WORKDIR)

clean:
	rm -rf "$(OUTPUT_DIR)"