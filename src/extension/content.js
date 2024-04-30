fetch('http://localhost:8000/api/v1/sites')
  .then(response => response.json())
  .then(data => {
    console.log(data)
    chrome.runtime.sendMessage({ action: 'updateSupportedSites', sites: data });
  })
  .catch(error => {
    console.log('Error while getting websites list:', error);
  });