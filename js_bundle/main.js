import HelloWorld from "./test_module";
import { createEditor } from "lexical";

HelloWorld();

const config = {
  namespace: "MyEditor",
  onError: console.error,
};

const editor = createEditor(config);

const editorRoot = document.getElementById("lexical-editor");

editor.setRootElement(editoRoot);
