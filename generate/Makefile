OUTPUT_DIR ?= "../fsdata"

obj = $(shell seq 500 500 25000) $(shell seq 50000 25000 500000)

all: build

build: $(addprefix build/,$(obj))

$(addprefix build/,$(obj)):
	$(eval WORKDIR := $(shell mktemp -d))
	mkdir -p "$(WORKDIR)" "$(OUTPUT_DIR)"
	truncate "$(WORKDIR)/$(subst build/,,$@)m" --size "$(subst build/,,$@)m"
	mkfs.ext4 -F "$(WORKDIR)/$(subst build/,,$@)m"
	go run . --input="$(WORKDIR)/$(subst build/,,$@)m" --size="$(subst build/,,$@)" --output-dir="$(OUTPUT_DIR)"
	rm -rf $(WORKDIR)

clean:
	rm -rf "$(OUTPUT_DIR)"