
inputSet = {"a1", "a2", "a2"}
inputMap1 = {"a1": "", "a2": ""}
inputList2 = ["a2", "a3", "a2"]
inputMap2 = {"a2": "", "a3": ""}
outputList1 = ["a1", "a2"]
outputList2 = ["a2", "a3"]
intersection = ["a2"]
union = ["a1", "a2", "a3"]
difference = ["a1"]
symmetric_difference = ["a1", "a3"]


print("inputSet", inputSet)
print("inputList2", inputList2)
s = set(inputList2)
print("set from inputList2", s)

print("inputMap1", inputMap1)
s = set(inputMap1)
print("set from inputMap1", s)

# union with list
print("union of inputSet and inputList2")
print("    using operator", inputSet | set(inputList2))
print("    using list", inputSet.union(inputList2))
print("    expected", union)

# union with map
print("union of inputSet and inputMap2")
print("    using operator", inputSet | set(inputMap2))
print("    using list", inputSet.union(inputMap2))
print("    expected", union)

# intersection with list
print("intersection of inputSet and inputList2")
print("    using operator", inputSet & set(inputList2))
print("    using list", inputSet.intersection(inputList2))
print("    expected", intersection)

# intersection with map
print("intersection of inputSet and inputMap2")
print("    using operator", inputSet & set(inputMap2))
print("    using list", inputSet.intersection(inputMap2))
print("    expected", intersection)

# difference with list
print("difference of inputSet and inputList2")
print("    using operator", inputSet - set(inputList2))
print("    using list", inputSet.difference(inputList2))
print("    expected", difference)

# difference with map
print("difference of inputSet and inputMap2")
print("    using operator", inputSet - set(inputMap2))
print("    using list", inputSet.difference(inputMap2))
print("    expected", difference)

# symmetric difference with list
print("symmetric difference of inputSet and inputList2")
print("    using operator", inputSet ^ set(inputList2))
print("    using list", inputSet.symmetric_difference(inputList2))
print("    expected", symmetric_difference)

# symmetric difference with map
print("symmetric difference of inputSet and inputMap2")
print("    using operator", inputSet ^ set(inputMap2))
print("    using list", inputSet.symmetric_difference(inputMap2))
print("    expected", symmetric_difference)
