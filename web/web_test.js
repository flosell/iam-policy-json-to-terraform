import { Selector } from 'testcafe'; // first import testcafe selectors

fixture `iam-policy-json-to-terraform web version`.page `./index.html`;  // specify the start page


//then create a test and place your code there
test('happy path', async t => {
    await t
        .click('#doConvert')
        .expect(Selector('#output').value).eql('data "aws_iam_policy_document" "hello" {}\n');
});

test('error case', async t => {
    await t
        .selectText('#input')
        .pressKey('delete')
        .typeText('#input', '{')
        .click('#doConvert')
        .expect(Selector('#output').value).contains('unexpected end of JSON input');
});