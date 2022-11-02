package main

const NoOfBuckets = 7

// A HashTable stores hashed sting keys using seperate chaining to avoid collisions.
type HashTable struct {
	buckets [NoOfBuckets]*SinglyLinkedList[string]
}

// Insert inserts the key into the Hashtable.
func (ht *HashTable) Insert(key string) {
	bucket := ht.getBucket(ht.toHash(key))
	// Storing duplicates here but could have an "if !bucket.Search(key)" guard.
	// Have decided against that as it would require traversing the bucket on every Insert
	// and any duplicates would be removed on bucket.Delete(value) anyway.
	bucket.Prepend(key)
}

// Search returns whether the key is stored in the HashTable.
func (ht *HashTable) Search(key string) bool {
	bucket := ht.buckets[ht.toHash(key)]
	if bucket == nil {
		return false
	}
	return bucket.Search(key)
}

// Delete removes the key from the HashTable
func (ht *HashTable) Delete(key string) {
	hash := ht.toHash(key)
	bucket := ht.buckets[hash]
	if bucket == nil {
		return
	}

	bucket.Delete(key)

	if bucket.length == 0 {
		ht.buckets[hash] = nil
	}
}

func (ht *HashTable) getBucket(hash rune) *SinglyLinkedList[string] {
	if ht.buckets[hash] == nil {
		ht.buckets[hash] = &SinglyLinkedList[string]{}
	}
	return ht.buckets[hash]
}

func (ht *HashTable) toHash(key string) (hash rune) {
	for _, r := range key {
		hash += r
	}
	return hash % NoOfBuckets
}
