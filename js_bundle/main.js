import HelloWorld from "./test_module";
import { createEditor } from "lexical";

HelloWorld();

const config = {
  namespace: "MyEditor",
  onError: console.error,
};

const editor = createEditor(config);
