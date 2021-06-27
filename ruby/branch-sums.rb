def branch_sums(root)
    return [] if root.nil?

    result = []
    collect_sum(root, 0, result)
    return result
end

def collect_sum(node, acc, result)
    return acc if node.nil?

    acc_latest = acc + node.value

    unless(node.left.nil? && node.right.nil?)
        result.push(acc_latest)
    end

    collect_sum(node.left, acc_latest, result)
    collect_sum(node.right, acc_latest, result)
end