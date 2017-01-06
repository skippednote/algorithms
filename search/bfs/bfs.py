from collections import deque

def person_is_seller(name):
    return name[-1] == 'm'

graph = {}
graph["you"] = ["john", "rob", "ross"]
graph["jim"] = ["cat", "sim"]
graph["rob"] = []
graph["ross"] = ["suster"]

def search(name):
    seacrch_queue = deque()
    seacrch_queue += graph[name]
    searched = [] # store the nodes that have been already looked up.
    
    while seacrch_queue:
        person = seacrch_queue.popleft()
        if not person in searched:
            if person_is_seller(person):
                print(person + " is a mango seller.")
                return True
            else:
                seacrch_queue += graph[person]
                searched.append(person)
    return False

search("you")
