#!/bin/bash

# Script to add retry mechanism to all resource files in a directory
# Usage: ./add_retry_to_all_resources.sh <directory_path>

if [ $# -eq 0 ]; then
    echo "Usage: $0 <directory_path>"
    echo "Example: $0 internal/service"
    echo "Example: $0 internal/service/dns"
    exit 1
fi

SEARCH_PATH="$1"

if [[ ! -d "$SEARCH_PATH" ]]; then
    echo "Error: Directory '$SEARCH_PATH' not found!"
    exit 1
fi

echo "Adding retry mechanism to all resource files in: $SEARCH_PATH"
echo "================================================================"

# Find all *_resource.go files recursively
RESOURCE_FILES=$(find "$SEARCH_PATH" -name "*_resource.go" -type f)

if [[ -z "$RESOURCE_FILES" ]]; then
    echo "No resource files found in $SEARCH_PATH"
    exit 1
fi

echo "Found resource files:"
echo "$RESOURCE_FILES" | while read -r file; do
    echo "  • $file"
done
echo ""

# Counters
PROCESSED=0
SKIPPED=0
FAILED=0

# Function to check if file already has retry mechanism
has_retry_mechanism() {
    local file="$1"
    grep -q "retryOperation" "$file" && grep -q "time\.Duration" "$file"
}

# Function to check if import exists
has_import() {
    local file="$1"
    local import_path="$2"
    grep -q "\"$import_path\"" "$file"
}

# Function to add import if not exists
add_import() {
    local file="$1"
    local import_path="$2"
    if ! has_import "$file" "$import_path"; then
        # Find the last import line and add after it
        sed -i '' '/^import (/,/^)$/{
            /^)$/{
                i\
    "'"$import_path"'"
            }
        }' "$file"
        return 0
    fi
    return 1
}

# Function to get resource type name from file
get_resource_type() {
    local file="$1"
    grep -o "type .*Resource struct" "$file" | sed 's/type \(.*\)Resource struct.*/\1/' | head -1
}

# Function to add retry mechanism to a single file
add_retry_to_file() {
    local file="$1"
    echo "Processing: $(basename "$file")"
    
    # Check if already has retry mechanism
    if has_retry_mechanism "$file"; then
        echo "  ⏭  Already has retry mechanism, skipping"
        return 2
    fi
    
    # Create backup
    cp "$file" "$file.backup"
    
    local changes_made=0
    local resource_type
    resource_type=$(get_resource_type "$file")
    
    if [[ -z "$resource_type" ]]; then
        echo "  ❌ Could not determine resource type"
        return 1
    fi
    
    echo "  📝 Resource type: ${resource_type}Resource"
    
    # Step 1: Add required imports
    echo "  🔧 Adding imports..."
    local import_changes=0
    
    if add_import "$file" "time"; then
        ((import_changes++))
    fi
    if add_import "$file" "errors"; then
        ((import_changes++))
    fi
    if add_import "$file" "github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"; then
        ((import_changes++))
    fi
    
    if [[ $import_changes -gt 0 ]]; then
        echo "    ✓ Added $import_changes imports"
        changes_made=1
    else
        echo "    - All imports already exist"
    fi
    
    # Step 2: Add timeout variables to CRUD methods
    echo "  🔧 Adding variables to CRUD methods..."
    local method_changes=0
    
    # Add to Create method
    if grep -q "func (r \*${resource_type}Resource) Create(" "$file"; then
        if ! grep -A 5 "func (r \*${resource_type}Resource) Create(" "$file" | grep -q "var timeout time.Duration"; then
            sed -i '' "/func (r \*${resource_type}Resource) Create(/,/var data ${resource_type}Model/{
                /var data ${resource_type}Model/a\\
    var timeout time.Duration\\
    var httpRes *http.Response
            }" "$file"
            ((method_changes++))
        fi
    fi
    
    # Add to Read method
    if grep -q "func (r \*${resource_type}Resource) Read(" "$file"; then
        if ! grep -A 5 "func (r \*${resource_type}Resource) Read(" "$file" | grep -q "var timeout time.Duration"; then
            sed -i '' "/func (r \*${resource_type}Resource) Read(/,/var data ${resource_type}Model/{
                /var data ${resource_type}Model/a\\
    var timeout time.Duration\\
    var httpRes *http.Response
            }" "$file"
            ((method_changes++))
        fi
    fi
    
    # Add to Update method
    if grep -q "func (r \*${resource_type}Resource) Update(" "$file"; then
        if ! grep -A 5 "func (r \*${resource_type}Resource) Update(" "$file" | grep -q "var timeout time.Duration"; then
            sed -i '' "/func (r \*${resource_type}Resource) Update(/,/var data ${resource_type}Model/{
                /var data ${resource_type}Model/a\\
    var timeout time.Duration\\
    var httpRes *http.Response
            }" "$file"
            ((method_changes++))
        fi
    fi
    
    # Add to Delete method
    if grep -q "func (r \*${resource_type}Resource) Delete(" "$file"; then
        if ! grep -A 5 "func (r \*${resource_type}Resource) Delete(" "$file" | grep -q "var timeout time.Duration"; then
            sed -i '' "/func (r \*${resource_type}Resource) Delete(/,/var data ${resource_type}Model/{
                /var data ${resource_type}Model/a\\
    var timeout time.Duration\\
    var httpRes *http.Response
            }" "$file"
            ((method_changes++))
        fi
    fi
    
    if [[ $method_changes -gt 0 ]]; then
        echo "    ✓ Added variables to $method_changes methods"
        changes_made=1
    fi
    
    # Step 3: Add timeout initialization
    echo "  🔧 Adding timeout initialization..."
    if ! grep -q "TimeInSeconds.*ValueInt64" "$file"; then
        # Add timeout initialization after error check in each method
        sed -i '' '/if resp\.Diagnostics\.HasError() {/,/return/{
            /return/i\
\
    // Set timeout from the resource if available\
    if !data.TimeInSeconds.IsNull() && !data.TimeInSeconds.IsUnknown() {\
        timeout = time.Duration(data.TimeInSeconds.ValueInt64()) * time.Second\
    }
        }' "$file"
        echo "    ✓ Added timeout initialization"
        changes_made=1
    fi
    
    # Step 4: Wrap API calls with retry logic
    echo "  🔧 Wrapping API calls with retry logic..."
    
    # Pattern 1: Create method - apiRes, _, err := 
    sed -i '' '/apiRes, _, err := r\.client\./,/Execute()/{
        /apiRes, _, err := r\.client\./{
            s/apiRes, _, err := r\.client\./var apiRes */
            a\
\
    err := r.retryOperation(ctx, timeout, func() error {\
        var err error\
        apiRes, httpRes, err = r.client.
        }
        /Execute()/{
            a\
        return err\
    })
        }
    }' "$file"
    
    # Pattern 2: Read method - apiRes, httpRes, err :=
    sed -i '' '/apiRes, httpRes, err := r\.client\./,/Execute()/{
        /apiRes, httpRes, err := r\.client\./{
            s/apiRes, httpRes, err := r\.client\./var apiRes */
            a\
\
    err := r.retryOperation(ctx, timeout, func() error {\
        var err error\
        apiRes, httpRes, err = r.client.
        }
        /Execute()/{
            a\
        return err\
    })
        }
    }' "$file"
    
    # Pattern 3: Delete method - httpRes, err :=
    sed -i '' '/httpRes, err := r\.client\./,/Execute()/{
        /httpRes, err := r\.client\./{
            a\
\
    err := r.retryOperation(ctx, timeout, func() error {\
        var err error\
        httpRes, err = r.client.
        }
        /Execute()/{
            a\
        return err\
    })
        }
    }' "$file"
    
    echo "    ✓ Wrapped API calls with retry logic"
    changes_made=1
    
    # Step 5: Add ReadByExtAttrs timeout if method exists
    if grep -q "func (r \*${resource_type}Resource) ReadByExtAttrs(" "$file"; then
        echo "  🔧 Adding timeout to ReadByExtAttrs method..."
        
        # Add timeout variable to ReadByExtAttrs
        if ! grep -A 10 "func (r \*${resource_type}Resource) ReadByExtAttrs(" "$file" | grep -q "var timeout time.Duration"; then
            sed -i '' "/func (r \*${resource_type}Resource) ReadByExtAttrs(/,/var diags diag\.Diagnostics/{
                /var diags diag\.Diagnostics/a\\
    var timeout time.Duration
            }" "$file"
        fi
        
        # Add timeout initialization in ReadByExtAttrs
        sed -i '' '/if data\.ExtAttrsAll\.IsNull() {/i\
    // Set timeout from the resource if available\
    if !data.TimeInSeconds.IsNull() && !data.TimeInSeconds.IsUnknown() {\
        timeout = time.Duration(data.TimeInSeconds.ValueInt64()) * time.Second\
    }\
' "$file"
        
        # Wrap ReadByExtAttrs API call
        sed -i '' '/func (r \*'"${resource_type}"'Resource) ReadByExtAttrs(/,/^}$/{
            /apiRes, _, err := r\.client\./,/Execute()/{
                /apiRes, _, err := r\.client\./{
                    s/apiRes, _, err := r\.client\./var apiRes */
                    a\
\
    err := r.retryOperation(ctx, timeout, func() error {\
        var err error\
        apiRes, _, err = r.client.
                }
                /Execute()/{
                    a\
        return err\
    })
                }
            }
        }' "$file"
        
        echo "    ✓ Added timeout support to ReadByExtAttrs"
    fi
    
    # Step 6: Add retry helper methods
    echo "  🔧 Adding retry helper methods..."
    
    # Check if retryOperation already exists
    if ! grep -q "func (r \*${resource_type}Resource) retryOperation(" "$file"; then
        cat >> "$file" << EOF

func (r *${resource_type}Resource) retryOperation(
    ctx context.Context,
    timeout time.Duration,
    operation func() error,
) error {

    // If timeout is not set, execute once and return
    if timeout <= 0 {
        return operation()
    }

    return retry.RetryContext(ctx, timeout, func() *retry.RetryError {
        err := operation()

        // ---- SUCCESS ----
        if err == nil {
            return nil
        }

        // ---- CONTEXT / TERRAFORM TIMEOUT HANDLING ----
        // Never retry these, otherwise plugin can hang
        if errors.Is(err, context.DeadlineExceeded) ||
            errors.Is(ctx.Err(), context.DeadlineExceeded) ||
            errors.Is(ctx.Err(), context.Canceled) {

            return retry.NonRetryableError(
                fmt.Errorf("operation stopped due to Terraform timeout or cancellation: %w", err),
            )
        }

        // ---- PLACEHOLDER FOR FUTURE RETRY LOGIC ----
        // Currently disabled intentionally
        if isRetryableErrorPlaceholder(err) {
            return retry.RetryableError(err)
        }

        // ---- DEFAULT: FAIL FAST ----
        return retry.NonRetryableError(err)
    })
}

func isRetryableErrorPlaceholder(err error) bool {
    // IMPORTANT:
    // This function is intentionally conservative.
    // It always returns false today.
    //
    // Purpose:
    // - Acts as a safe extension point
    // - Allows retry logic to be added later
    // - Prevents accidental infinite retries

    return false
}
EOF
        echo "    ✓ Added retry helper methods"
        changes_made=1
    fi
    
    # Step 7: Clean up and validate
    echo "  🔧 Cleaning up..."
    
    # Remove duplicate variable declarations
    sed -i '' '/var timeout time\.Duration/{
        N
        /\n.*var timeout time\.Duration/d
    }' "$file"
    
    sed -i '' '/var httpRes \*http\.Response/{
        N
        /\n.*var httpRes \*http\.Response/d
    }' "$file"
    
    # Format the file if go is available
    if command -v go >/dev/null 2>&1; then
        if go fmt "$file" > /dev/null 2>&1; then
            echo "    ✓ Go formatting successful"
        else
            echo "    ⚠ Go formatting had issues (file may still be valid)"
        fi
    fi
    
    if [[ $changes_made -eq 1 ]]; then
        echo "  ✅ Successfully added retry mechanism"
        return 0
    else
        echo "  ⚠ No changes were needed"
        rm -f "$file.backup"
        return 2
    fi
}

# Process each file
echo "Processing files..."
echo ""

echo "$RESOURCE_FILES" | while read -r file; do
    if [[ -n "$file" ]]; then
        case $(add_retry_to_file "$file") in
            0)
                ((PROCESSED++))
                echo ""
                ;;
            1)
                ((FAILED++))
                echo "  ❌ Failed to process file"
                echo ""
                ;;
            2)
                ((SKIPPED++))
                echo ""
                ;;
        esac
    fi
done

# Final summary (note: counters won't work in subshell, so we'll count files differently)
TOTAL_FILES=$(echo "$RESOURCE_FILES" | wc -l)
PROCESSED_FILES=$(find "$SEARCH_PATH" -name "*_resource.go.backup" -type f | wc -l)
SKIPPED_FILES=$(echo "$RESOURCE_FILES" | xargs -I {} sh -c 'if grep -q "retryOperation" "{}"; then echo "{}"; fi' | wc -l)
FAILED_FILES=$((TOTAL_FILES - PROCESSED_FILES - SKIPPED_FILES))

echo "================================================================"
echo "🎉 Batch processing completed!"
echo ""
echo "📊 Summary:"
echo "  Total files found: $TOTAL_FILES"
echo "  Successfully processed: $PROCESSED_FILES"
echo "  Already had retry mechanism: $SKIPPED_FILES" 
echo "  Failed to process: $FAILED_FILES"
echo ""

if [[ $PROCESSED_FILES -gt 0 ]]; then
    echo "✅ Changes made to $PROCESSED_FILES files:"
    echo "  • Added required imports (time, errors, retry)"
    echo "  • Added timeout and httpRes variables to CRUD methods"
    echo "  • Added timeout initialization from TimeInSeconds field"
    echo "  • Wrapped all API calls with retry logic"
    echo "  • Added retry helper methods"
    echo ""
    echo "📁 Backup files created with .backup extension"
    echo ""
    echo "⚠️  Next steps:"
    echo "  1. Add 'time_in_seconds' field to resource schemas if not present"
    echo "  2. Test all modified resources thoroughly"
    echo "  3. Customize retry logic in isRetryableErrorPlaceholder() if needed"
    echo "  4. Remove backup files if everything works correctly:"
    echo "     find $SEARCH_PATH -name \"*.backup\" -delete"
fi

echo ""
echo "🔍 To see what files were modified:"
echo "    find $SEARCH_PATH -name \"*_resource.go.backup\""
echo ""
echo "🧹 To remove all backup files:"
echo "    find $SEARCH_PATH -name \"*_resource.go.backup\" -delete"