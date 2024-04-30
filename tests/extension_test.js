const sitesCompativeis = require('./sitesCompativeis');

describe('getSites', () => {

    it('should return true when argument is true', () => {
        expect(sitesCompativeis(true)).toBe(true);
    });

    it('should return true when argument is false', () => {
        expect(sitesCompativeis(false)).toBe(false);
    });
});
