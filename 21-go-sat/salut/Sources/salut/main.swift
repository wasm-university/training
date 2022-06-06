import Suborbital
import Foundation

class Salut: Suborbital.Runnable {
    func run(input: String) -> String {

        let words = input.components(separatedBy: [";"])


        return "Swift " + words.joined(separator: "-")

    }
}

Suborbital.Set(runnable: Salut())
