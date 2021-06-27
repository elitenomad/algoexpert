# Set the current node with the tree
# Assume the closest value is tree.value to start with
# Loop until the bst is over (top down approach)
# Update the closest value based on (target - closest) > target - current_node.value
# Have to go through right or left of the tree as the rule of BST is that
# value will be greater than the left side of the tree , less or equal on the right side
def find_closest_value_bst(tree, target)
    current_node = tree
    closest = tree.value

    while !current_node.nil?
        if((target - closest) > (target - current_node.value))
            closest = current_node.value
        end

        if target < current_node.value
            current_node = current_node.left
        elsif target > current_node.value
            current_node = current_node.right
        else
            break
        end
    end

    closest
end