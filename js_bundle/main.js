import HelloWorld from "./test_module";
import { createEditor } from "lexical";
import InitSignupPage from "./signup_page/signup_page";

HelloWorld();

const config = {
  namespace: "MyEditor",
  onError: console.error,
};

const editor = createEditor(config);

const editorRoot = document.getElementById("lexical-editor");

editor.setRootElement(editorRoot);
console.log(editor);

InitSignupPage();
