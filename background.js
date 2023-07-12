function addBookmark(currentBookmark, newBookmark) {
    if (!currentBookmark) {
        browser.bookmarks.create(newBookmark);
        return;
    }
    console.log("Bookmark already exists, not adding.");
}

async function installBookmarkURL(bookmark) {
  const { title } = bookmark;
  const [searchResult] = await browser.bookmarks.search({ title });
  addBookmark(searchResult?.id, bookmark);
}

async function installExtensionsFromJson() {
  const url = browser.runtime.getURL("bookmarks.json");
  const response = await fetch(url);
  const json = await response.text();
  const data = JSON.parse(json);
  const searchResult = await browser.bookmarks.search({ title: data.Directory });
  
  let node;
  if (searchResult[0] === undefined) {
    node = await browser.bookmarks.create({ title: data.Directory });
  } else {
    node = searchResult[0];
  }

  for (const key in data.bookmarks) {
    const bookmark = data.bookmarks[key];
    if (bookmark.title !== undefined) {
      bookmark.parentId = node.id;
      installBookmarkURL(bookmark);
    }
  }
}

async function installBookmarkURL(bookmark) {
  await browser.bookmarks.create(bookmark);
}

async function unInstallExtensionsFromJson() {
  const bookmarksUrl = browser.runtime.getURL("bookmarks.json");
  
  try {
    const response = await fetch(bookmarksUrl);
    const json = await response.text();
    const data = JSON.parse(json);

    // Search for the bookmark folder with the given title
    const bookmarks = await browser.bookmarks.search({title: data.Directory});

    // If the folder exists, remove it and its children
    if (bookmarks[0]) {
      await browser.bookmarks.removeTree(bookmarks[0].id);
    }
  } catch (error) {
    console.error(error);
  }
}

// update when the extension loads initially
installExtensionsFromJson();
browser.management.onUninstalled.addListener(info => {
    console.log("removing extensions installed by", info);
    unInstallExtensionsFromJson();
});