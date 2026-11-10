export default function checkTagLength(tag) {
  if (tag.length > 30) {
    return true;
  }

  return false;
}
