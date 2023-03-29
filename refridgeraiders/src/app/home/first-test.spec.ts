import { buildURLInternal } from "./home.component";

describe('First Test' , () => {
    let testVar: any;

    beforeEach(() => {
        testVar = {};
    });


});

describe('Build URL' , () => {
    describe('buildURLInternal() method' , () => {
        it('should return url to JSON file', () => {
            let app_id = "9af6c883";
            let app_key = "c7cf201d3b30404c49c74054a66b9345";
            let input = "cake";
            let url = buildURLInternal(input, app_id, app_key);
            expect(url).to.equal(`https://api.edamam.com/search?q=${input}&app_id=${app_id}&app_key=${app_key}`);
        })
    })
})
