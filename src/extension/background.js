chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
    if (request.action === 'updateSupportedSites') {
      const supportedSites = request.sites;
      
      if (supportedSites.includes(window.location.href)) {
        console.log('it is a supported website mf' + window.location.href);
      }
    }
  });