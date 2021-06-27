class TwoSum
    attr_accessor :input, :target
    def initialize(input:, target:)
        @input = input
        @target = target
    end

    def v1
        return [] if input.length <= 0
        
        result = []
        input.each do |elem|
            difference = input.index(target - elem).to_i
            if difference > 0
                result =  [elem, input[difference]]
            end
        end

        return result
    end

    def v2
        return [] if input.length <= 0

        result = {}

        input.each do |elem|
            prospect = target - elem
            if result[prospect]
                return [elem, prospect]
            end

            result[elem] = true
        end

        return []
    end
end