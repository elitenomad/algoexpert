# Depth First search algorithm
class Node
    def initialize(name)
        @name = name
        @children = []
    end

    def add_child(name)
        @children << new Node(name)
        self
    end

    # Empty array as input
    def dpfs(array = [])
        # Push the first element
        array << @name

        @children.each do |child|
            child.dpfs(array)
        end

        return array
    end
end