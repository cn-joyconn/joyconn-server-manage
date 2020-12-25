package cache
import(
	cache2go "github.com/muesli/cache2go"
)

//ICache struct
type  ICache struct {
	catalog      string
	cacheName string 
}
func (iCache  *ICache)getkey(key string) (string){
	return cacheName+"\f"+key;
}
//Put a cache obj
func (iCache  *ICache)Put(key string, val string ,seconds int) error{
	cache := cache2go.Cache(iCache.catalog);
	key = getkey(key)
	cache2go.Add(key, seconds*time.Second, &val)
	cache.Flush()
	return nil;
}

//Remove a cache key
func (iCache  *ICache)Remove(key string) error{
	cache := cache2go.Cache(iCache.catalog);
	key = getkey(key)
	cache2go.Delete(key)
	cache.Flush()
	return nil;
}

//Get a cached obj
func (iCache  *ICache)Get(key string) (string,error){
	cache := cache2go.Cache(iCache.catalog);
	key = getkey(key)
	res,err:=cache2go.Value(key)
	return &res,err
}