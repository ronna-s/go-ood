package repo

//
//// Copyright 2016 The Go Authors. All rights reserved.
//// Use of this source code is governed by a BSD-style
//// license that can be found in the LICENSE file.
//import (
//	"sync"
//	"sync/atomic"
//	"unsafe"
//)
//
//// Map is like a Go map[interface{}]interface{} but is safe for concurrent use
//// by multiple goroutines without additional locking or coordination.
//// Loads, stores, and deletes run in amortized constant time.
////
//// The Map type is specialized. Most code should use a plain Go map instead,
//// with separate locking or coordination, for better type safety and to make it
//// easier to maintain other invariants along with the map content.
////
//// The Map type is optimized for two common use cases: (1) when the entry for a given
//// key is only ever written once but read many times, as in caches that only grow,
//// or (2) when multiple goroutines read, write, and overwrite entries for disjoint
//// sets of keys. In these two cases, use of a Map may significantly reduce lock
//// contention compared to a Go map paired with a separate Mutex or RWMutex.
////
//// The zero Map is empty and ready for use. A Map must not be copied after first use.
////
//// In the terminology of the Go memory model, Map arranges that a write operation
//// “synchronizes before” any read operation that observes the effect of the write, where
//// read and write operations are defined as follows.
//// Load, LoadAndDelete, LoadOrStore are read operations;
//// Delete, LoadAndDelete, and Store are write operations;
//// and LoadOrStore is a write operation when it returns loaded set to false.
//type Map[K comparable, V interface{}] struct {
//	mu sync.Mutex
//
//	// read contains the portion of the map's contents that are safe for
//	// concurrent access (with or without mu held).
//	//
//	// The read field itself is always safe to load, but must only be stored with
//	// mu held.
//	//
//	// Entries stored in read may be updated concurrently without mu, but updating
//	// a previously-expunged entry requires that the entry be copied to the dirty
//	// map and unexpunged with mu held.
//	read atomic.Value // readOnly
//
//	// dirty contains the portion of the map's contents that require mu to be
//	// held. To ensure that the dirty map can be promoted to the read map quickly,
//	// it also includes all of the non-expunged entries in the read map.
//	//
//	// Expunged entries are not stored in the dirty map. An expunged entry in the
//	// clean map must be unexpunged and added to the dirty map before a new value
//	// can be stored to it.
//	//
//	// If the dirty map is nil, the next write to the map will initialize it by
//	// making a shallow copy of the clean map, omitting stale entries.
//	dirty map[K]*entry
//
//	// misses counts the number of loads since the read map was last updated that
//	// needed to lock mu to determine whether the key was present.
//	//
//	// Once enough misses have occurred to cover the cost of copying the dirty
//	// map, the dirty map will be promoted to the read map (in the unamended
//	// state) and the next store to the map will make a new dirty copy.
//	misses int
//
//	// expunged is an arbitrary pointer that marks entries which have been deleted
//	// from the dirty map.
//	expunged unsafe.Pointer
//}
//
//func NewMap[K comparable, V interface{}]() *Map[K, V] {
//	return &Map[K, V]{expunged: unsafe.Pointer(new(V))}
//
//}
//
//// readOnly is an immutable struct stored atomically in the Map.read field.
//type readOnly[K comparable, V interface{}] struct {
//	m       map[K]*entry[V]
//	amended bool // true if the dirty map contains some key not in m.
//}
//
//// An entry is a slot in the map corresponding to a particular key.
//type entry[V interface{}] struct {
//	// p points to the interface{} value stored for the entry.
//	//
//	// If p == nil, the entry has been deleted, and either m.dirty == nil or
//	// m.dirty[key] is e.
//	//
//	// If p == expunged, the entry has been deleted, m.dirty != nil, and the entry
//	// is missing from m.dirty.
//	//
//	// Otherwise, the entry is valid and recorded in m.read.m[key] and, if m.dirty
//	// != nil, in m.dirty[key].
//	//
//	// An entry can be deleted by atomic replacement with nil: when m.dirty is
//	// next created, it will atomically replace nil with expunged and leave
//	// m.dirty[key] unset.
//	//
//	// An entry's associated value can be updated by atomic replacement, provided
//	// p != expunged. If p == expunged, an entry's associated value can be updated
//	// only after first setting m.dirty[key] = e so that lookups using the dirty
//	// map find the entry.
//	p V
//}
//
//func newEntry[V interface{}](i V) *entry[V] {
//	return &entry[V]{p: i}
//}
//
//// Load returns the value stored in the map for a key, or nil if no
//// value is present.
//// The ok result indicates whether value was found in the map.
//func (m *Map) Load(key any) (value any, ok bool) {
//	read, _ := m.read.Load().(readOnly)
//	e, ok := read.m[key]
//	if !ok && read.amended {
//		m.mu.Lock()
//		// Avoid reporting a spurious miss if m.dirty got promoted while we were
//		// blocked on m.mu. (If further loads of the same key will not miss, it's
//		// not worth copying the dirty map for this key.)
//		read, _ = m.read.Load().(readOnly)
//		e, ok = read.m[key]
//		if !ok && read.amended {
//			e, ok = m.dirty[key]
//			// Regardless of whether the entry was present, record a miss: this key
//			// will take the slow path until the dirty map is promoted to the read
//			// map.
//			m.missLocked()
//		}
//		m.mu.Unlock()
//	}
//	if !ok {
//		return nil, false
//	}
//	return e.load()
//}
//
//func (e *entry) load() (value any, ok bool) {
//	p := atomic.LoadPointer(&e.p)
//	if p == nil || p == expunged {
//		return nil, false
//	}
//	return *(*any)(p), true
//}
//
//// Store sets the value for a key.
//func (m *Map) Store(key, value any) {
//	read, _ := m.read.Load().(readOnly)
//	if e, ok := read.m[key]; ok && e.tryStore(&value) {
//		return
//	}
//
//	m.mu.Lock()
//	read, _ = m.read.Load().(readOnly)
//	if e, ok := read.m[key]; ok {
//		if e.unexpungeLocked() {
//			// The entry was previously expunged, which implies that there is a
//			// non-nil dirty map and this entry is not in it.
//			m.dirty[key] = e
//		}
//		e.storeLocked(&value)
//	} else if e, ok := m.dirty[key]; ok {
//		e.storeLocked(&value)
//	} else {
//		if !read.amended {
//			// We're adding the first new key to the dirty map.
//			// Make sure it is allocated and mark the read-only map as incomplete.
//			m.dirtyLocked()
//			m.read.Store(readOnly{m: read.m, amended: true})
//		}
//		m.dirty[key] = newEntry(value)
//	}
//	m.mu.Unlock()
//}
//
//// tryStore stores a value if the entry has not been expunged.
////
//// If the entry is expunged, tryStore returns false and leaves the entry
//// unchanged.
//func (e *entry) tryStore(i *any) bool {
//	for {
//		p := atomic.LoadPointer(&e.p)
//		if p == expunged {
//			return false
//		}
//		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
//			return true
//		}
//	}
//}
//
//// unexpungeLocked ensures that the entry is not marked as expunged.
////
//// If the entry was previously expunged, it must be added to the dirty map
//// before m.mu is unlocked.
//func (e *entry) unexpungeLocked() (wasExpunged bool) {
//	return atomic.CompareAndSwapPointer(&e.p, expunged, nil)
//}
//
//// storeLocked unconditionally stores a value to the entry.
////
//// The entry must be known not to be expunged.
//func (e *entry) storeLocked(i *any) {
//	atomic.StorePointer(&e.p, unsafe.Pointer(i))
//}
//
//// LoadOrStore returns the existing value for the key if present.
//// Otherwise, it stores and returns the given value.
//// The loaded result is true if the value was loaded, false if stored.
//func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool) {
//	// Avoid locking if it's a clean hit.
//	read, _ := m.read.Load().(readOnly)
//	if e, ok := read.m[key]; ok {
//		actual, loaded, ok := e.tryLoadOrStore(value)
//		if ok {
//			return actual, loaded
//		}
//	}
//
//	m.mu.Lock()
//	read, _ = m.read.Load().(readOnly)
//	if e, ok := read.m[key]; ok {
//		if e.unexpungeLocked() {
//			m.dirty[key] = e
//		}
//		actual, loaded, _ = e.tryLoadOrStore(value)
//	} else if e, ok := m.dirty[key]; ok {
//		actual, loaded, _ = e.tryLoadOrStore(value)
//		m.missLocked()
//	} else {
//		if !read.amended {
//			// We're adding the first new key to the dirty map.
//			// Make sure it is allocated and mark the read-only map as incomplete.
//			m.dirtyLocked()
//			m.read.Store(readOnly{m: read.m, amended: true})
//		}
//		m.dirty[key] = newEntry(value)
//		actual, loaded = value, false
//	}
//	m.mu.Unlock()
//
//	return actual, loaded
//}
//
//// tryLoadOrStore atomically loads or stores a value if the entry is not
//// expunged.
////
//// If the entry is expunged, tryLoadOrStore leaves the entry unchanged and
//// returns with ok==false.
//func (e *entry) tryLoadOrStore(i any) (actual any, loaded, ok bool) {
//	p := atomic.LoadPointer(&e.p)
//	if p == expunged {
//		return nil, false, false
//	}
//	if p != nil {
//		return *(*any)(p), true, true
//	}
//
//	// Copy the interface after the first load to make this method more amenable
//	// to escape analysis: if we hit the "load" path or the entry is expunged, we
//	// shouldn't bother heap-allocating.
//	ic := i
//	for {
//		if atomic.CompareAndSwapPointer(&e.p, nil, unsafe.Pointer(&ic)) {
//			return i, false, true
//		}
//		p = atomic.LoadPointer(&e.p)
//		if p == expunged {
//			return nil, false, false
//		}
//		if p != nil {
//			return *(*any)(p), true, true
//		}
//	}
//}
//
//// LoadAndDelete deletes the value for a key, returning the previous value if any.
//// The loaded result reports whether the key was present.
//func (m *Map) LoadAndDelete(key any) (value any, loaded bool) {
//	read, _ := m.read.Load().(readOnly)
//	e, ok := read.m[key]
//	if !ok && read.amended {
//		m.mu.Lock()
//		read, _ = m.read.Load().(readOnly)
//		e, ok = read.m[key]
//		if !ok && read.amended {
//			e, ok = m.dirty[key]
//			delete(m.dirty, key)
//			// Regardless of whether the entry was present, record a miss: this key
//			// will take the slow path until the dirty map is promoted to the read
//			// map.
//			m.missLocked()
//		}
//		m.mu.Unlock()
//	}
//	if ok {
//		return e.delete()
//	}
//	return nil, false
//}
//
//// Delete deletes the value for a key.
//func (m *Map) Delete(key any) {
//	m.LoadAndDelete(key)
//}
//
//func (e *entry) delete() (value any, ok bool) {
//	for {
//		p := atomic.LoadPointer(&e.p)
//		if p == nil || p == expunged {
//			return nil, false
//		}
//		if atomic.CompareAndSwapPointer(&e.p, p, nil) {
//			return *(*any)(p), true
//		}
//	}
//}
//
//// Range calls f sequentially for each key and value present in the map.
//// If f returns false, range stops the iteration.
////
//// Range does not necessarily correspond to any consistent snapshot of the Map's
//// contents: no key will be visited more than once, but if the value for any key
//// is stored or deleted concurrently (including by f), Range may reflect any
//// mapping for that key from any point during the Range call. Range does not
//// block other methods on the receiver; even f itself may call any method on m.
////
//// Range may be O(N) with the number of elements in the map even if f returns
//// false after a constant number of calls.
//func (m *Map) Range(f func(key, value any) bool) {
//	// We need to be able to iterate over all of the keys that were already
//	// present at the start of the call to Range.
//	// If read.amended is false, then read.m satisfies that property without
//	// requiring us to hold m.mu for a long time.
//	read, _ := m.read.Load().(readOnly)
//	if read.amended {
//		// m.dirty contains keys not in read.m. Fortunately, Range is already O(N)
//		// (assuming the caller does not break out early), so a call to Range
//		// amortizes an entire copy of the map: we can promote the dirty copy
//		// immediately!
//		m.mu.Lock()
//		read, _ = m.read.Load().(readOnly)
//		if read.amended {
//			read = readOnly{m: m.dirty}
//			m.read.Store(read)
//			m.dirty = nil
//			m.misses = 0
//		}
//		m.mu.Unlock()
//	}
//
//	for k, e := range read.m {
//		v, ok := e.load()
//		if !ok {
//			continue
//		}
//		if !f(k, v) {
//			break
//		}
//	}
//}
//
//func (m *Map) missLocked() {
//	m.misses++
//	if m.misses < len(m.dirty) {
//		return
//	}
//	m.read.Store(readOnly{m: m.dirty})
//	m.dirty = nil
//	m.misses = 0
//}
//
//func (m *Map) dirtyLocked() {
//	if m.dirty != nil {
//		return
//	}
//
//	read, _ := m.read.Load().(readOnly)
//	m.dirty = make(map[any]*entry, len(read.m))
//	for k, e := range read.m {
//		if !e.tryExpungeLocked() {
//			m.dirty[k] = e
//		}
//	}
//}
//
//func (e *entry) tryExpungeLocked() (isExpunged bool) {
//	p := atomic.LoadPointer(&e.p)
//	for p == nil {
//		if atomic.CompareAndSwapPointer(&e.p, nil, expunged) {
//			return true
//		}
//		p = atomic.LoadPointer(&e.p)
//	}
//	return p == expunged
//}
