def node_depths_count_on(tree)
    return recurring_depths(tree, 0)
end

def recurring_depths(root, acc)
    return acc + recurring_depths(root.left, acc) + recurring_depths(root.right, acc) 
end

def node_depths_count_on_v2(tree)
    sum_of_depths = 0
    stack = [{
        node: tree,
        depth: 0
    }]

    while(stack.length > 0)
        obj = stack.pop
        if obj.nil 
            continue
        end

        sum_of_depths += obj.depth

        stack << [{ node: tree.left, depth: depth + 1}]
        stack << [{ node: tree.right, depth: depth + 1}]
    end

    return sum_of_depths
end

class BinaryTree 
    def initialize(value)
        @value = value
        @left = nil
        @right = nil
    end
end