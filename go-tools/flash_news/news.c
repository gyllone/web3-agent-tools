#include <tools.h>
#include <news.h>

IMPL_LIST(NewsItem)

void release_NewsItem(NewsItem item) {
	release_String(item.title);
	release_String(item.content);
	release_Int(item.unixtime);
}

void release_NewsResult(NewsResult res) {
	release_List_NewsItem(res.items);
	release_Bool(res.success);
}

void release_NewsResultJson(NewsResultJson res) {
	release_String(res.json_string);
	release_Bool(res.success);
}