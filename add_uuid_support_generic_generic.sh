#!/bin/bash

# Universal script to add UUID support to Terraform provider model and data source test files
# Usage: ./add_uuid_support_universal.sh [input_file]
# Input file format: pairs of model_*.go and *_data_source_test.go files

INPUT_FILE="${1:-m_d_testfiles.txt}"

if [[ ! -f "$INPUT_FILE" ]]; then
    echo "Error: Input file '$INPUT_FILE' not found!"
    exit 1
fi

echo "Adding UUID support to files from $INPUT_FILE..."
echo "=================================================="

# Universal function to extract struct name from any model file
get_struct_name_universal() {
    local model_file="$1"
    
    # Try to find the struct definition in the file and extract the name
    local struct_line=$(grep -m 1 "^type .*Model struct" "$model_file" 2>/dev/null)
    
    if [[ -n "$struct_line" ]]; then
        # Extract struct name from "type StructNameModel struct"
        echo "$struct_line" | sed 's/type \(.*\)Model struct.*/\1/'
    else
        # Fallback: generate from filename
        local base=$(basename "$model_file" .go)
        base="${base#model_}"
        # Convert to PascalCase
        echo "$base" | sed 's/_/ /g' | awk '{for(i=1;i<=NF;i++){$i=toupper(substr($i,1,1)) tolower(substr($i,2))}} 1' | sed 's/ //g'
    fi
}

# Universal function to find test function pattern
get_test_function_universal() {
    local test_file="$1"
    
    # Look for ResourceAttrPair function in the test file
    local func_line=$(grep -m 1 "func.*ResourceAttrPair" "$test_file" 2>/dev/null)
    
    if [[ -n "$func_line" ]]; then
        # Extract function name
        echo "$func_line" | sed 's/func \([^(]*\).*/\1/'
    else
        # Fallback: generate from filename
        local base=$(basename "$test_file" _data_source_test.go)
        local converted=$(echo "$base" | sed 's/_/ /g' | awk '{for(i=1;i<=NF;i++){$i=toupper(substr($i,1,1)) tolower(substr($i,2))}} 1' | sed 's/ //g')
        echo "testAccCheck${converted}ResourceAttrPair"
    fi
}

# Universal function to add UUID support to any model file
add_uuid_to_model_universal() {
    local model_file="$1"
    local struct_name="$2"
    
    echo "Processing model file: $model_file"
    echo "Detected struct name: $struct_name"
    
    # Verify struct exists
    if ! grep -q "type ${struct_name}Model struct" "$model_file"; then
        echo "  ✗ Struct definition not found: type ${struct_name}Model struct"
        return 1
    fi
    
    echo "  ✓ Found struct definition"
    
    # 1. Add Uuid field to Model struct
    if ! grep -q 'Uuid.*types\.String.*`tfsdk:"uuid"`' "$model_file"; then
        echo "  Adding Uuid field to struct..."
        
        # Find the Ref field and add Uuid after it
        sed -i '' "/type ${struct_name}Model struct/,/^}/ {
            /Ref[[:space:]]*types\.String[[:space:]]*\`tfsdk:\"ref\"\`/a\\
    Uuid        types.String \`tfsdk:\"uuid\"\`
        }" "$model_file"
        
        if grep -q 'Uuid.*types\.String.*`tfsdk:"uuid"`' "$model_file"; then
            echo "  ✓ Added Uuid field to ${struct_name}Model struct"
        else
            echo "  ✗ Failed to add Uuid field to struct"
        fi
    else
        echo "  - Uuid field already exists"
    fi
    
    # 2. Add uuid to AttrTypes map
    if ! grep -q '"uuid".*types\.StringType' "$model_file"; then
        echo "  Adding uuid to AttrTypes map..."
        
        # Find the AttrTypes map and add uuid
        sed -i '' "/var ${struct_name}AttrTypes.*map/,/^}/ {
            /\"ref\":[[:space:]]*types\.StringType,/a\\
    \"uuid\":        types.StringType,
        }" "$model_file"
        
        if grep -q '"uuid".*types\.StringType' "$model_file"; then
            echo "  ✓ Added uuid to ${struct_name}AttrTypes"
        else
            echo "  ✗ Failed to add uuid to AttrTypes"
        fi
    else
        echo "  - uuid already exists in AttrTypes"
    fi
    
    # 3. Add uuid schema attribute (if ResourceSchemaAttributes exists)
    if grep -q "var ${struct_name}ResourceSchemaAttributes" "$model_file"; then
        if ! grep -q '"uuid": schema\.StringAttribute' "$model_file"; then
            echo "  Adding uuid schema attribute..."
            
            sed -i '' "/var ${struct_name}ResourceSchemaAttributes.*map/,/^}/ {
                /\"ref\": schema\.StringAttribute{/,/^[[:space:]]*},*[[:space:]]*$/ {
                    /^[[:space:]]*},*[[:space:]]*$/a\\
    \"uuid\": schema.StringAttribute{\\
        Computed:            true,\\
        MarkdownDescription: \"The uuid to the object.\",\\
    },
                }
            }" "$model_file"
            
            if grep -q '"uuid": schema\.StringAttribute' "$model_file"; then
                echo "  ✓ Added uuid schema attribute"
            else
                echo "  ✗ Failed to add uuid schema attribute"
            fi
        else
            echo "  - uuid schema attribute already exists"
        fi
    else
        echo "  - No ResourceSchemaAttributes found (skipping schema)"
    fi
    
    # 4. Add Uuid flattening in Flatten method (if it exists)
    if grep -q "func (m \*${struct_name}Model) Flatten" "$model_file"; then
        if ! grep -q 'm\.Uuid = flex\.FlattenStringPointer(from\.Uuid)' "$model_file"; then
            echo "  Adding Uuid flattening..."
            
            sed -i '' "/func (m \*${struct_name}Model) Flatten/,/^}/ {
                /m\.Ref = flex\.FlattenStringPointer(from\.Ref)/a\\
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
            }" "$model_file"
            
            if grep -q 'm\.Uuid = flex\.FlattenStringPointer(from\.Uuid)' "$model_file"; then
                echo "  ✓ Added Uuid flattening to Flatten method"
            else
                echo "  ✗ Failed to add Uuid flattening"
            fi
        else
            echo "  - Uuid flattening already exists"
        fi
    else
        echo "  - No Flatten method found (skipping flatten)"
    fi
}

# Universal function to add UUID assertion to any test file
add_uuid_to_test_universal() {
    local test_file="$1"
    local function_name="$2"
    
    echo "Processing test file: $test_file"
    echo "Detected function name: $function_name"
    
    # Check if the function exists
    if ! grep -q "func ${function_name}" "$test_file"; then
        echo "  ✗ Test function not found: $function_name"
        return 1
    fi
    
    echo "  ✓ Found test function"
    
    # Add UUID assertion if not already present
    if ! grep -q 'resource\.TestCheckResourceAttrPair.*"uuid".*"result\.0\.uuid"' "$test_file"; then
        echo "  Adding UUID assertion..."
        
        sed -i '' "/func ${function_name}/,/^}/ {
            /resource\.TestCheckResourceAttrPair.*\"ref\".*\"result\.0\.ref\"/a\\
        resource.TestCheckResourceAttrPair(resourceName, \"uuid\", dataSourceName, \"result.0.uuid\"),
        }" "$test_file"
        
        if grep -q 'resource\.TestCheckResourceAttrPair.*"uuid".*"result\.0\.uuid"' "$test_file"; then
            echo "  ✓ Added UUID assertion to test function"
        else
            echo "  ✗ Failed to add UUID assertion"
        fi
    else
        echo "  - UUID assertion already exists"
    fi
}

# Main processing function
process_file_pair() {
    local model_file="$1"
    local test_file="$2"
    
    # Find actual file paths
    local model_path=$(find . -name "$model_file" -type f 2>/dev/null | head -1)
    local test_path=$(find . -name "$test_file" -type f 2>/dev/null | head -1)
    
    if [[ ! -f "$model_path" ]]; then
        echo "  ✗ Model file not found: $model_file"
        return 1
    fi
    
    if [[ ! -f "$test_path" ]]; then
        echo "  ✗ Test file not found: $test_file"
        return 1
    fi
    
    # Extract names using universal functions
    local struct_name=$(get_struct_name_universal "$model_path")
    local test_function=$(get_test_function_universal "$test_path")
    
    echo "  Model: $model_path (Struct: $struct_name)"
    echo "  Test:  $test_path (Function: $test_function)"
    
    # Process both files
    add_uuid_to_model_universal "$model_path" "$struct_name"
    add_uuid_to_test_universal "$test_path" "$test_function"
    
    return 0
}

# Main processing loop
processed_count=0
error_count=0

echo "Reading file pairs from $INPUT_FILE..."
echo ""

while IFS= read -r line; do
    # Skip empty lines and comments
    [[ -z "$line" || "$line" =~ ^[[:space:]]*# ]] && continue
    
    # Remove leading ! if present
    line="${line#!}"
    
    # Check if this is a model file
    if [[ "$line" =~ ^model_.*\.go$ ]]; then
        model_file="$line"
        
        # Read the next line for the corresponding data source test file
        if read -r next_line; then
            next_line="${next_line#!}"
            
            # Check if next line is a test file
            if [[ "$next_line" =~ _data_source_test\.go$ ]]; then
                test_file="$next_line"
                
                echo "Processing pair $((++processed_count)):"
                
                if process_file_pair "$model_file" "$test_file"; then
                    echo "  ✓ Successfully processed pair"
                else
                    echo "  ✗ Failed to process pair"
                    ((error_count++))
                fi
                
                echo ""
            else
                echo "Warning: Expected test file after $model_file, but got: $next_line"
                echo ""
            fi
        else
            echo "Warning: No test file found after $model_file"
            echo ""
        fi
    fi
done < "$INPUT_FILE"

# Final summary
echo "=================================================="
echo "Processing Summary:"
echo "  Total pairs processed: $processed_count"
echo "  Successful: $((processed_count - error_count))"
echo "  Failed: $error_count"
echo ""

if [[ $processed_count -gt 0 ]]; then
    echo "UUID support has been added to all available files!"
    echo ""
    echo "What was added:"
    echo "  ✓ Uuid field to model structs"
    echo "  ✓ uuid entry to AttrTypes maps"
    echo "  ✓ uuid schema attributes (where applicable)"
    echo "  ✓ Uuid flattening to Flatten methods (where applicable)"
    echo "  ✓ UUID assertions to test functions"
    echo ""
    echo "Please review the changes and run tests to ensure everything works correctly."
else
    echo "No files were processed. Please check:"
    echo "  - Input file format and content"
    echo "  - File paths and locations"
    echo "  - File permissions"
fi