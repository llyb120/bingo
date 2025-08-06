#!/bin/bash

# update_requires.sh
# Usage: ./update_requires.sh <new_version>

if [ -z "$1" ]; then
    echo "Usage: $0 <new_version>"
    echo "Example: $0 v0.0.1"
    exit 1
fi

NEW_VERSION=$1
REPO_PREFIX="github.com/llyb120/bingo"

# Declare an array to hold tags
declare -a TAGS

# Function to update a single go.mod file
update_go_mod() {
    local modfile=$1
    local mod_name=$(grep "^module " "$modfile" | awk '{print $2}')
    
    echo "Updating $modfile"
    
    # Create a temporary file
    local tmpfile=$(mktemp)
    
    # Track if we're in require section
    local in_require=false
    
    while IFS= read -r line; do
        # Check for require section start/end
        if [[ "$line" =~ ^require[[:space:]]*\( ]]; then
            in_require=true
            echo "$line" >> "$tmpfile"
            continue
        elif [[ "$line" =~ ^\) ]] && [ "$in_require" = true ]; then
            in_require=false
            echo "$line" >> "$tmpfile"
            continue
        fi
        
        # Process require lines
        if [ "$in_require" = true ] || [[ "$line" =~ ^require[[:space:]] ]]; then
            # Check if this line contains our repo prefix
            if [[ "$line" == *"$REPO_PREFIX"* ]]; then
                # Extract the module name using awk
                local mod=$(echo "$line" | awk -v prefix="$REPO_PREFIX" '{
                    for(i=1;i<=NF;i++) {
                        if($i ~ prefix) {
                            print $i
                            break
                        }
                    }
                }')
                
                # Skip if this is the module itself
                if [[ "$mod" != "$mod_name" ]]; then
                    # Update version while preserving format using awk
                    local new_line=$(echo "$line" | awk -v mod="$mod" -v ver="$NEW_VERSION" '{
                        for(i=1;i<=NF;i++) {
                            if($i == mod) {
                                $(i+1) = ver
                            }
                        }
                        print
                    }')
                    echo "$new_line" >> "$tmpfile"
                    continue
                fi
            fi
        fi
        
        # Copy other lines as-is
        echo "$line" >> "$tmpfile"
    done < "$modfile"
    
    # Replace the original file
    mv "$tmpfile" "$modfile"

    # Add tag name to array for later creation
    TAGS+=("${mod_name}/${NEW_VERSION}")
}

# Find all go.mod files
find . -name "go.mod" | while read -r modfile; do
    update_go_mod "$modfile"
done

echo "All requires updated to $NEW_VERSION"
echo "Running tidy.sh..."
bash tidy.sh

echo "Staging changes..."
git add .

echo "Committing changes..."
git commit -m "chore: Update module dependencies to ${NEW_VERSION}"

echo "Creating tags..."
for tag in "${TAGS[@]}"; do
    echo "Creating tag: $tag"
    git tag "$tag"
done

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
echo "Pushing commits to origin/$CURRENT_BRANCH..."
git push origin "$CURRENT_BRANCH"

echo "Pushing all tags..."
git push --tags