import { Selector, t } from 'testcafe';

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

class Page {
    constructor() {
        this.output = Selector('#output')
        this.input = Selector('#input')
        this.convertButton = Selector('#doConvert')
    }

    async replaceInputText(newText) {
        await t
            .selectText(this.input)
            .pressKey('delete')
            .typeText(this.input, newText)
    }
}

let p = new Page()

test('happy path', async t => {
    await t.expect(p.output.value)
        .eql('data "aws_iam_policy_document" "hello" {}\n')

    await p.replaceInputText(someIamJson)
    await t.click(p.convertButton)

    await t.expect(p.output.value)
        .eql(someIamTerraform)
});

test('error case', async t => {
    await p.replaceInputText('{')
    await t.click(p.convertButton)

    await t.expect(p.output.value)
        .contains('unexpected end of JSON input')
});

test.page`./index.html#content=${encodeURIComponent(someIamJson)}`('bookmarklets', async t => {
    await t
        .expect(p.output.value).eql(someIamTerraform);
});