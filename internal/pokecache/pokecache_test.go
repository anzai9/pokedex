package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
  const interval = 5 * time.Second
  cases := []struct{
    key string
    val []byte
  }{
    {
      key: "https://example.com",
      val: []byte("testdata"),
    },
    {
      key: "https://example.com/2",
      val: []byte("testdata2"),
    },
  }

  for i, c := range cases {
    t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
      cache := NewCache(interval)
      cache.Add(c.key, c.val)
      val, ok := cache.Get(c.key)
      if !ok {
	t.Errorf("expected to find key")
	return
      }
      if string(val) != string(c.val) {
	t.Errorf("expected to find value")
	return
      }
    })
  }
}

func TestReapLoop(t *testing.T) {
  const baseTime = 5 * time.Millisecond
  const waitTime = baseTime + 5*time.Millisecond
  cache := NewCache(baseTime)
  key := "https://example.com"
  cache.Add(key, []byte("testdata"))

  if _, exists := cache.Get(key); !exists {
    t.Errorf("expected to find key")
    return
  }

  time.Sleep(waitTime)

  if _, exists := cache.Get(key); exists {
    t.Errorf("expected key to be reaped")
    return
  }
}
