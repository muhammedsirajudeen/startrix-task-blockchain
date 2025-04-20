"use client"

import type React from "react"
import { type Dispatch, type SetStateAction, useEffect, useState } from "react"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogClose,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import axiosInstance from "@/lib/axiosInstance"
import { toast } from "sonner"
import { ArrowRight, Database, FileSignature, Hash } from "lucide-react"

interface BlockExplorerProps {
  open: boolean
  setOpen: Dispatch<SetStateAction<boolean>>
}

interface Transaction {
  sender: string
  recipient: string
  amount: number
  signature: string
  previous_block: string
}

const BlockExplorer: React.FC<BlockExplorerProps> = ({ open, setOpen }) => {
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function getBlocks() {
      setLoading(true)
      try {
        const response = await axiosInstance.get("/transactions")
        const transactions: Transaction[] = response.data
        setTransactions(transactions)
      } catch (error) {
        console.log(error)
        toast.error(<p className="text-white">Unable to load the blocks</p>, { style: { backgroundColor: "red" } })
      } finally {
        setLoading(false)
      }
    }

    if (open) {
      getBlocks()
    }
  }, [open])

  // Function to truncate long strings like addresses and hashes
  const truncateString = (str: string, start = 6, end = 4) => {
    if (!str) return ""
    if (str.length <= start + end) return str
    return `${str.slice(0, start)}...${str.slice(-end)}`
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100 max-w-3xl max-h-[80vh] overflow-hidden flex flex-col">
        <DialogHeader>
          <DialogTitle className="text-xl font-bold flex items-center gap-2">
            <div className="bg-gradient-to-r from-amber-400 via-yellow-500 to-amber-600 rounded-full p-2">
              <Database className="h-5 w-5 text-white" />
            </div>
            Block Explorer
          </DialogTitle>
          <DialogDescription className="text-sm text-gray-600 dark:text-gray-400">
            Explore the details of blockchain transactions in real-time.
          </DialogDescription>
        </DialogHeader>

        <div className="mt-4 overflow-y-auto flex-1 pr-2">
          {loading ? (
            <div className="flex justify-center items-center h-40">
              <div className="animate-pulse flex flex-col items-center">
                <div className="bg-gradient-to-r from-amber-400 via-yellow-500 to-amber-600 rounded-full p-3">
                  <Database className="h-6 w-6 text-white" />
                </div>
                <p className="mt-2 text-gray-500">Loading transactions...</p>
              </div>
            </div>
          ) : transactions.length === 0 ? (
            <div className="text-center py-10">
              <p className="text-gray-500">No transactions found</p>
            </div>
          ) : (
            <div className="space-y-4">
              {transactions.map((tx, index) => (
                <div
                  key={index}
                  className="border border-gray-200 dark:border-gray-700 rounded-lg overflow-hidden shadow-sm hover:shadow-md transition-shadow"
                >
                  <div className="bg-gradient-to-r from-amber-400 via-yellow-500 to-amber-600 p-3 text-white flex justify-between items-center">
                    <div className="flex items-center gap-2">
                      <Hash className="h-4 w-4" />
                      <span className="font-medium">Transaction #{index + 1}</span>
                    </div>
                    <Badge variant="outline" className="bg-white/20 text-white border-none">
                      {tx.amount} COIN
                    </Badge>
                  </div>

                  <div className="p-4 bg-white dark:bg-gray-800 space-y-3">
                    <div className="flex flex-col sm:flex-row sm:items-center gap-2 sm:gap-4">
                      <div className="text-sm font-medium text-gray-500 dark:text-gray-400 w-24">From:</div>
                      <div className="flex-1 font-mono text-sm bg-gray-100 dark:bg-gray-700 p-2 rounded overflow-hidden overflow-ellipsis">
                        {tx.sender}
                      </div>
                    </div>

                    <div className="flex justify-center">
                      <ArrowRight className="text-amber-500" />
                    </div>

                    <div className="flex flex-col sm:flex-row sm:items-center gap-2 sm:gap-4">
                      <div className="text-sm font-medium text-gray-500 dark:text-gray-400 w-24">To:</div>
                      <div className="flex-1 font-mono text-sm bg-gray-100 dark:bg-gray-700 p-2 rounded overflow-hidden overflow-ellipsis">
                        {tx.recipient}
                      </div>
                    </div>


                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
                      <div>
                        <div className="text-xs text-gray-500 dark:text-gray-400 flex items-center gap-1">
                          <FileSignature className="h-3 w-3" /> Signature
                        </div>
                        <div className="font-mono text-xs mt-1 truncate">{truncateString(tx.signature, 10, 10)}</div>
                      </div>
                      <div>
                        <div className="text-xs text-gray-500 dark:text-gray-400 flex items-center gap-1">
                          <Database className="h-3 w-3" /> Previous Block
                        </div>
                        <div className="font-mono text-xs mt-1 truncate">
                          {truncateString(tx.previous_block, 10, 10)}
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>

        <div className="mt-4 pt-2 border-t border-gray-200 dark:border-gray-700 flex justify-between items-center">
          <div className="text-sm text-gray-500">
            {transactions.length} transaction{transactions.length !== 1 ? "s" : ""} found
          </div>
          <DialogClose asChild>
            <Button variant="secondary">Close</Button>
          </DialogClose>
        </div>
      </DialogContent>
    </Dialog>
  )
}

export default BlockExplorer
