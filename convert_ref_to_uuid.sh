#!/bin/bash

# Enhanced script to convert API calls from ref to uuid in resource files
# Usage: ./convert_ref_to_uuid_enhanced.sh [file1] [file2] ... or ./convert_ref_to_uuid_enhanced.sh *.go

if [ $# -eq 0 ]; then
    echo "Usage: $0 <file1> [file2] ... or $0 *.go"
    echo "Example: $0 record_aaaa_resource.go"
    echo "Example: $0 *_resource.go"
    exit 1
fi

echo "Converting API calls from ref to uuid (Enhanced)..."
echo "================================================="

processed_count=0
error_count=0

# Function to convert a single file
convert_file_enhanced() {
    local file="$1"
    
    if [[ ! -f "$file" ]]; then
        echo "  ✗ File not found: $file"
        return 1
    fi
    
    if [[ ! "$file" =~ _resource\.go$ ]]; then
        echo "  - Skipping non-resource file: $file"
        return 0
    fi
    
    echo "Processing: $file"
    
    # Create backup
    cp "$file" "$file.backup"
    
    # Track if any changes were made
    local changes_made=0
    
    # 1. Convert Read method API calls - more specific pattern
    if grep -q "\.Read(ctx, utils\.ExtractResourceRef(data\.Ref\.ValueString()))" "$file"; then
        sed -i '' 's/\.Read(ctx, utils\.ExtractResourceRef(data\.Ref\.ValueString()))/\.Read(ctx, data.Uuid.ValueString())/g' "$file"
        echo "  ✓ Updated Read method to use uuid"
        changes_made=1
    fi
    
    # 2. Convert Update method API calls - more specific pattern  
    if grep -q "\.Update(ctx, utils\.ExtractResourceRef(data\.Ref\.ValueString()))" "$file"; then
        sed -i '' 's/\.Update(ctx, utils\.ExtractResourceRef(data\.Ref\.ValueString()))/\.Update(ctx, data.Uuid.ValueString())/g' "$file"
        echo "  ✓ Updated Update method to use uuid"
        changes_made=1
    fi
    
    # 3. Convert Delete method API calls - more specific pattern
    if grep -q "\.Delete(ctx, utils\.ExtractResourceRef(data\.Ref\.ValueString()))" "$file"; then
        sed -i '' 's/\.Delete(ctx, utils\.ExtractResourceRef(data\.Ref\.ValueString()))/\.Delete(ctx, data.Uuid.ValueString())/g' "$file"
        echo "  ✓ Updated Delete method to use uuid"
        changes_made=1
    fi
    
    # 4. Handle multiline patterns where the method call might be split
    # Pattern: Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
    if grep -A1 -B1 "utils\.ExtractResourceRef(data\.Ref\.ValueString())" "$file" | grep -q "\.Read(ctx"; then
        sed -i '' '/\.Read(ctx,/{
            N
            s/\.Read(ctx, utils\.ExtractResourceRef(data\.Ref\.ValueString()))/\.Read(ctx, data.Uuid.ValueString())/g
        }' "$file"
        echo "  ✓ Updated multiline Read method to use uuid"
        changes_made=1
    fi
    
    # 5. Handle multiline Update patterns
    if grep -A1 -B1 "utils\.ExtractResourceRef(data\.Ref\.ValueString())" "$file" | grep -q "\.Update(ctx"; then
        sed -i '' '/\.Update(ctx,/{
            N
            s/\.Update(ctx, utils\.ExtractResourceRef(data\.Ref\.ValueString()))/\.Update(ctx, data.Uuid.ValueString())/g
        }' "$file"
        echo "  ✓ Updated multiline Update method to use uuid"
        changes_made=1
    fi
    
    # 6. Handle multiline Delete patterns
    if grep -A1 -B1 "utils\.ExtractResourceRef(data\.Ref\.ValueString())" "$file" | grep -q "\.Delete(ctx"; then
        sed -i '' '/\.Delete(ctx,/{
            N
            s/\.Delete(ctx, utils\.ExtractResourceRef(data\.Ref\.ValueString()))/\.Delete(ctx, data.Uuid.ValueString())/g
        }' "$file"
        echo "  ✓ Updated multiline Delete method to use uuid"
        changes_made=1
    fi
    
    # 7. Direct line-by-line replacement for remaining utils.ExtractResourceRef calls
    if grep -q "utils\.ExtractResourceRef(data\.Ref\.ValueString())" "$file"; then
        echo "  ✓ Converting remaining utils.ExtractResourceRef calls..."
        sed -i '' 's/utils\.ExtractResourceRef(data\.Ref\.ValueString())/data.Uuid.ValueString()/g' "$file"
        changes_made=1
    fi
    
    # 8. Update state attribute references in Update method (if not already done)
    if grep -q 'req\.State\.GetAttribute(ctx, path\.Root("ref"), &data\.Ref)' "$file"; then
        sed -i '' 's/req\.State\.GetAttribute(ctx, path\.Root("ref"), &data\.Ref)/req.State.GetAttribute(ctx, path.Root("uuid"), \&data.Uuid)/g' "$file"
        echo "  ✓ Updated state attribute reference to use uuid"
        changes_made=1
    fi
    
    # 9. Update ImportState method to set uuid instead of ref (if not already done)
    if grep -q 'resp\.State\.SetAttribute(ctx, path\.Root("ref"), req\.ID)' "$file"; then
        sed -i '' 's/resp\.State\.SetAttribute(ctx, path\.Root("ref"), req\.ID)/resp.State.SetAttribute(ctx, path.Root("uuid"), req.ID)/g' "$file"
        echo "  ✓ Updated ImportState to set uuid"
        changes_made=1
    fi
    
    # 10. Check if all utils.ExtractResourceRef calls are removed
    if grep -q "utils\.ExtractResourceRef" "$file"; then
        echo "  ⚠ Warning: Still found remaining utils.ExtractResourceRef usage in $file"
        echo "    Remaining occurrences:"
        grep -n "utils\.ExtractResourceRef" "$file" || true
        echo "    Attempting final cleanup..."
        
        # Final aggressive cleanup - replace any remaining utils.ExtractResourceRef patterns
        sed -i '' 's/utils\.ExtractResourceRef([^)]*)/data.Uuid.ValueString()/g' "$file"
        
        if ! grep -q "utils\.ExtractResourceRef" "$file"; then
            echo "  ✓ Successfully cleaned up all utils.ExtractResourceRef references"
            changes_made=1
        fi
    else
        echo "  ✓ All utils.ExtractResourceRef references successfully converted"
    fi
    
    # 11. Check for associate_internal_id logic
    if grep -q "associate_internal_id" "$file"; then
        echo "  ⚠ Info: Found associate_internal_id logic in $file"
        echo "    Note: UUID-based resources typically don't need associate_internal_id logic"
        echo "    You may want to review and potentially remove this logic"
    fi
    
    # 12. Verify the conversion was successful
    if ! grep -q "utils\.ExtractResourceRef" "$file"; then
        echo "  ✅ Verification: No utils.ExtractResourceRef calls remaining"
    else
        echo "  ❌ Verification failed: Some utils.ExtractResourceRef calls still exist"
        grep -n "utils\.ExtractResourceRef" "$file"
    fi
    
    if [[ $changes_made -eq 1 ]]; then
        echo "  ✓ Successfully converted $file from ref to uuid"
        return 0
    else
        echo "  - No changes needed in $file"
        # Restore from backup since no changes were needed
        mv "$file.backup" "$file"
        return 0
    fi
}

# Process all provided files
for file in "$@"; do
    echo ""
    if convert_file_enhanced "$file"; then
        ((processed_count++))
    else
        ((error_count++))
    fi
done

echo ""
echo "================================================="
echo "Enhanced Conversion Summary:"
echo "  Files processed: $processed_count"
echo "  Errors: $error_count"
echo ""

if [[ $processed_count -gt 0 ]]; then
    echo "✅ Enhanced conversion completed!"
    echo ""
    echo "Changes made:"
    echo "  • All API calls now use data.Uuid.ValueString() instead of utils.ExtractResourceRef(data.Ref.ValueString())"
    echo "  • State attributes reference 'uuid' instead of 'ref'"
    echo "  • ImportState sets 'uuid' instead of 'ref'"
    echo "  • Removed all utils.ExtractResourceRef usage"
    echo ""
    echo "Backup files created with .backup extension"
    echo ""
    echo "Next steps:"
    echo "  1. Review the associate_internal_id logic (may not be needed for UUID-based resources)"
    echo "  2. Test the resource operations thoroughly"
    echo "  3. Remove backup files if everything works correctly"
    
    # Show files that were actually modified
    echo ""
    echo "Modified files:"
    for file in "$@"; do
        if [[ -f "$file.backup" ]]; then
            if ! cmp -s "$file" "$file.backup"; then
                echo "  • $file"
                echo "    Changes: $(grep -c "data.Uuid.ValueString()" "$file" || echo "0") uuid references added"
            fi
        fi
    done
else
    echo "No files were modified."
fi

# Cleanup empty backup files
for file in "$@"; do
    if [[ -f "$file.backup" ]]; then
        if cmp -s "$file" "$file.backup"; then
            rm "$file.backup"
        fi
    fi
done