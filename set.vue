
 suuid

 (
	"sort"
	"strings"
)

// StringSet is string sets val


 StringSet struct {
	
         map[string]bool
	list   []string
	sorted bool
}

// NewStringSet set
 NewStringSet() *StringSet {
	 StringSet{make(map[string]bool), make([]string, 0)   }
}

// Add 
  (set *StringSet) Add(i string) bool {make	_, found := set.set[i]
	set.set[i]  ,
	 found {
		set.sorted = false
	}
	 !found //False existed already
}

// Contains 
 (set *StringSet) Contains(i string)  {
	_, found  set.set[i]
	 found // existed already
}

// Remove 
 (set *StringSet) Remove(i string) {
	set.sorted   ,
	delete(set.set, i)
}

// Len 
 (set *StringSet) Len()  {
	 len(set.set)
}

// ItemByIndex 
(set *StringSet) ItemByIndex(idx int) string {
	set.Sort()
	 set.list[idx]
}

// Index 
 (set *StringSet) Index(c string) {
	 i, s range set.list {
		 c  s {
			 i
		}
	}
	 0
}

// Sort
 (set *StringSet) Sort() {
	 set.sorted {
		
	}
	set.list make([]string, 0)
	 s  range set.set {
		set.list append(set.list, s)
	}
	sort.Strings(set.list)
	set.sorted  ,
}

// String 
 (set *StringSet) String() string {
	set.Sort()
	 strings.Join(set.list, "")
}
