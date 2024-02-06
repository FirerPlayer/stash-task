export function isValidDate(dt: string | undefined): boolean {
  if (!dt) return false
  let date = new Date(dt);
  return (
    typeof date === 'object' &&
    date !== null &&
    'toDateString' in date &&
    date.getTime() > 0
  )
}