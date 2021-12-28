
golang sets
===========

Implementing sets in golang...

Sets are useful when implementing filtering. This is usually done 
by the intersection of a selection set and a target object. The intersection 
only returns the attributes of the target object specified in the selection set.

Using the sets clas from python3 as an example:

```python
# Filter a list
object = ["a1", "a3"]
selected = set("a1", "a2").intersection(object)

returns 
    {"a1"}

# Filter a map:
object = {"a1":"value1", "a3":"value3"}
for f in selected = set("a1", "a2").intersection(object):
    m[f] = object[f]
```

The opposite case is redact all fields that are not selected:

```python
selector = {"a1", "a2"}
object = {"a1":"value1", "a3":"value3"}
for f in deselected = set(object).difference(selector):
    object[f] = "redacted"

```

See sets.py for full list of examples.
