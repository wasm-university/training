import { logInfo } from "@suborbital/suborbital"

export function run(input: ArrayBuffer): ArrayBuffer {
	let inStr = String.UTF8.decode(input)
  
	let out = "AssemblyScript: " + inStr.split(";").join("-")

	logInfo(`👋 log from the runnable: ${inStr}`)

	return String.UTF8.encode(out)
}