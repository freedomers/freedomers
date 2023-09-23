def Merge(dict1, dict2): 
    res = {**dict1, **dict2} 
    return res 
      
# 两个字典
dict1 = {"name": "Joy", "age": 25}
dict2 = {"name": "Joy", "city": "New York"}
dict3 = Merge(dict1, dict2) 
print(dict3)
