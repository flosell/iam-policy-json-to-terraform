import { Selector } from 'testcafe'; // first import testcafe selectors

fixture `iam-policy-json-to-terraform web version`.page `./index.html`;  // specify the start page

let someIamJson = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "FirstStatement",
      "Effect": "Allow",
      "Action": ["iam:ChangePassword"],
      "Resource": "*"
    }
  ]
}`

let someIamTerraform = `data "aws_iam_policy_document" "hello" {
  statement {
    sid       = "FirstStatement"
    effect    = "Allow"
    resources = ["*"]
    actions   = ["iam:ChangePassword"]
  }
}
`

//then create a test and place your code there
test('happy path', async t => {
    await t
        // .click('#doConvert')
        .expect(Selector('#output').value).eql('data "aws_iam_policy_document" "hello" {}\n')
        .selectText('#input')
        .pressKey('delete')
        .typeText('#input', someIamJson)
        .click('#doConvert')
        .expect(Selector('#output').value).eql(someIamTerraform);
});

test('error case', async t => {
    await t
        .selectText('#input')
        .pressKey('delete')
        .typeText('#input', '{')
        .click('#doConvert')
        .expect(Selector('#output').value).contains('unexpected end of JSON input');
});

test.page`./index.html#content=${encodeURIComponent(someIamJson)}`('bookmarklets', async t => {
    await t
        .expect(Selector('#output').value).eql(someIamTerraform);
});