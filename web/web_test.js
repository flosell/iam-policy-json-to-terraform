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
        this.error = Selector('#json-error')
        this.infoToggle = Selector('#info-toggle')
        this.infoExpander = Selector('#info-expander')
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
    await t.expect(p.output.textContent)
        .eql('data "aws_iam_policy_document" "hello" {}\n')

    await p.replaceInputText(someIamJson)

    await t.expect(p.output.textContent)
        .eql(someIamTerraform)
});

test('error case', async t => {
    await p.replaceInputText('{')

    await t.expect(p.error.innerText)
        .contains('unexpected end of JSON input')
});

test('update to error case', async t => {
    await p.replaceInputText(someIamJson)

    await t.expect(p.output.textContent)
        .eql(someIamTerraform)

    await p.replaceInputText('{')

    await t.expect(p.output.textContent)
        .eql(someIamTerraform)

    await t.expect(p.error.innerText)
        .contains('unexpected end of JSON input')

    await p.replaceInputText('{}')

    await t.expect(p.error.innerText)
        .contains('input did not contain any statements')
});

test('collapse infos', async t => {
    await t.expect(p.infoExpander.visible).ok()

    await t
        .click(p.infoToggle)

    await t.expect(p.infoExpander.visible).notOk()
})

test('bookmarklets', async t => {
    await t.navigateTo(`./index.html#content=${encodeURIComponent(someIamJson)}`)
    await t
        .expect(p.output.textContent).eql(someIamTerraform);
});
