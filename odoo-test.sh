#!/bin/bash
set -e

# Default list of addons to test if no arguments are provided
DEFAULT_ADDONS="tifshop_product_product"

# Use the first command-line argument if it exists, otherwise use the default
ADDONS_TO_TEST="${1:-$DEFAULT_ADDONS}"

# --- Create include patterns for the coverage report from the addon list ---
OIFS=$IFS
IFS=','
COVERAGE_PATTERNS=""
for addon in $ADDONS_TO_TEST; do
    if [ -z "$COVERAGE_PATTERNS" ]; then
        COVERAGE_PATTERNS="*/$addon/*"
    else
        COVERAGE_PATTERNS="$COVERAGE_PATTERNS,*/$addon/*"
    fi
done
IFS=$OIFS
# ------------------------------------------------------------------------

COMPOSE_FILE="docker-compose.test.yaml"

echo "Building test image..."
docker-compose -f $COMPOSE_FILE build

echo "Installing and testing addons: $ADDONS_TO_TEST"

# Define the command to be executed inside the container
# The coverage report now uses --include to filter the output
TEST_COMMAND="coverage run --rcfile=.coveragerc -m odoo \
    -d test_db \
    --addons-path=./addons,./additional-addons \
    -i ${ADDONS_TO_TEST} \
    --test-enable \
    --stop-after-init \
    --log-level=test ; coverage report -m --include='${COVERAGE_PATTERNS}'"

# Execute the command using docker-compose run
docker-compose -f $COMPOSE_FILE run --rm odoo /bin/bash -c "$TEST_COMMAND"

echo "Test run finished. Cleaning up..."
docker-compose -f $COMPOSE_FILE down -v

echo "Done."
