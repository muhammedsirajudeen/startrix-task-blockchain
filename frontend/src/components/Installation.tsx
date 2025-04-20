"use client"

import { useState } from "react"
import { Check, Copy } from "lucide-react"
import { Button } from "@/components/ui/button"

export function CliInstall() {
  const [copied, setCopied] = useState(false)
  const installCommand =
    "bash <(wget -qO- https://raw.githubusercontent.com/muhammedsirajudeen/startrix-task-blockchain/refs/heads/main/cli/cli-install.sh)"

  const copyToClipboard = async () => {
    try {
      await navigator.clipboard.writeText(installCommand)
      setCopied(true)
      setTimeout(() => setCopied(false), 2000)
    } catch (err) {
      console.error("Failed to copy: ", err)
    }
  }

  return (
    <div className="rounded-xl w-96 overflow-hidden border border-gray-200 bg-gray-50 dark:bg-gray-900 dark:border-gray-800">
      <div className="flex items-center justify-between px-4 py-2 bg-gray-100 dark:bg-gray-800">
        <span className="text-sm font-medium">Install Cold Wallet CLI - Linux</span>
        <Button
          onClick={copyToClipboard}
          className="py-1 px-3 text-xs rounded-lg bg-gradient-to-r from-yellow-500 to-amber-500 hover:from-yellow-600 hover:to-amber-600 shadow-md hover:shadow-lg transition-all duration-300 hover:scale-[1.02]"
        >
          {copied ? <Check className="h-3.5 w-3.5 mr-1" /> : <Copy className="h-3.5 w-3.5 mr-1" />}
          {copied ? "Copied!" : "Copy"}
        </Button>
      </div>
      <div className="p-4 overflow-x-auto">
        <pre className="text-sm font-mono whitespace-pre-wrap break-all">{installCommand}</pre>
      </div>
    </div>
  )
}
