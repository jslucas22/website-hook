const supportedSites = [];

fetch('http://localhost:8000/api/v1/sitesWithLogin')
  .then(response => response.json())
  .then(data => {
    supportedSites.push(...data);
  })
  .catch(error => {
    console.log('err-01: Erro ao obter lista de sites:', error);
  });

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (message.action === 'updateSupportedSites') {
    supportedSites.push(...message.sites);
  }
});

setInterval(() => {
  const currentURL = window.location.href;

  if (supportedSites.some(site => site.url === currentURL)) {
    const userSiteData = supportedSites.find(site => site.url === currentURL);
    if (userSiteData) {
      switch (true) {
        case currentURL.includes("pernambucanas"):
          const plusoftOmniStrategy = new PlusoftOmniStrategy(userSiteData.username_input, userSiteData.password_input);
          plusoftOmniStrategy.login();
          break;
        default:
          // {...}
          break;
      }
    }
  }
}, 300);
