check_vegeta:
	@if  ! which vegeta; then \
    		echo "vegeta not found. Installing..."; \
    		brew install vegeta; \
	else \
    		echo "vegeta is already installed ✅"; \
    fi

attack: check_vegeta
	echo "GET http://localhost:8080/v1/code-blocks/?ids=1%2C2%2C3,4" | vegeta attack -duration=${duration} --rate=${rate} | vegeta report --type=${type}
