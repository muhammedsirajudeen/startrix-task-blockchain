"use client"

import type React from "react"

import { useState, useEffect } from "react"
import { Code, Clock, Copy, Coffee, CoffeeIcon, Coins } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Textarea } from "@/components/ui/textarea"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { toast, Toaster } from "sonner"

interface Transaction {
  amount: number
  recipient: string
  sender: string
  signature: string
  timestamp?: number
  status?: string
}

export default function TransactionPage() {
  const [jsonInput, setJsonInput] = useState("")
  const [transaction, setTransaction] = useState<Transaction | null>(null)
  const [history, setHistory] = useState<Transaction[]>([])
  const [isValid, setIsValid] = useState(false)

  useEffect(() => {
    const savedHistory = localStorage.getItem("transactionHistory")
    if (savedHistory) {
      setHistory(JSON.parse(savedHistory))
    }
  }, [])

  const handleInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const input = e.target.value
    setJsonInput(input)

    try {
      const parsed = JSON.parse(input)
      if (
        typeof parsed.amount === "number" &&
        typeof parsed.recipient === "string" &&
        typeof parsed.sender === "string" &&
        typeof parsed.signature === "string"
      ) {
        setTransaction(parsed)
        setIsValid(true)
      } else {
        setTransaction(null)
        setIsValid(false)
      }
    } catch (error) {
      console.log(error)
      setTransaction(null)
      setIsValid(false)
    }
  }

  const handleSubmit = () => {
    if (transaction) {
      setTimeout(() => {
        const completedTransaction = {
          ...transaction,
          timestamp: Date.now(),
          status: "completed",
        }

        const updatedHistory = [completedTransaction, ...history]
        setHistory(updatedHistory)

        localStorage.setItem("transactionHistory", JSON.stringify(updatedHistory))

        toast.success("Coffee Purchased!")

        setJsonInput("")
        setTransaction(null)
        setIsValid(false)
      }, 1500)

      toast.info("Pending")
    }
  }

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text)
    toast.success("Copied to clipboard")
  }

  // Format address for display
  const formatAddress = (address: string) => {
    if (!address) return ""
    return `${address.substring(0, 8)}...${address.substring(address.length - 8)}`
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-amber-50 to-white">
      <div className="container mx-auto px-4 py-12">
        <header className="mb-12 text-center">
          <div className="inline-block mb-4 p-2 bg-white rounded-full shadow-lg">
            <div className="bg-gradient-to-r from-amber-400 via-yellow-500 to-amber-600 rounded-full p-3">
              <Coffee className="h-8 w-8 text-white" />
            </div>
          </div>
          <h1 className="text-4xl font-bold mb-2 bg-clip-text text-transparent bg-gradient-to-r from-amber-600 via-yellow-600 to-amber-700">
            Coffee Coin
          </h1>
          <p className="text-amber-800/70">Support creators with secure blockchain transactions</p>
        </header>

        <Tabs defaultValue="submit" className="w-full max-w-4xl mx-auto">
          <TabsList className="grid w-full grid-cols-2 mb-8 p-1 bg-white rounded-xl shadow-md border border-amber-100">
            <TabsTrigger
              value="submit"
              className="text-lg rounded-lg data-[state=active]:bg-gradient-to-r data-[state=active]:from-amber-400 data-[state=active]:to-yellow-500 data-[state=active]:text-white"
            >
              <CoffeeIcon className="mr-2 h-5 w-5" />
              Buy Coffee
            </TabsTrigger>
            <TabsTrigger
              value="history"
              className="text-lg rounded-lg data-[state=active]:bg-gradient-to-r data-[state=active]:from-yellow-500 data-[state=active]:to-amber-500 data-[state=active]:text-white"
            >
              <Clock className="mr-2 h-5 w-5" />
              Coffee History
            </TabsTrigger>
          </TabsList>

          <TabsContent value="submit">
            <div className="grid gap-8 md:grid-cols-2">
              {/* JSON Input Card */}
              <Card className="bg-white border border-amber-100 shadow-xl rounded-2xl overflow-hidden transition-all duration-300 hover:shadow-2xl">
                <CardHeader className="text-center w-full">
                  <CardTitle className="flex items-center gap-2 text-xl">
                    <Code className="h-5 w-5" />
                    Transaction Details
                  </CardTitle>
                  <CardDescription className="text-white/80">Paste your transaction JSON below</CardDescription>
                </CardHeader>
                <CardContent className="pt-6">
                  <Textarea
                    value={jsonInput}
                    onChange={handleInputChange}
                    placeholder={`{\n  "amount": 100,\n  "recipient": "abcd",\n  "sender": "b38a7456c8d12db59760999a4c51f1e6b611b3f0953a0274c9dc92115eb30a95",\n  "signature": "pAc1SeETVXZOiHI0NnKrCbWhFzsGSNVSspZfNe7CcobKIhJuM5qyQhLUVmnY4mQQ6oKyDjd/q8cH84HraLCtAQ=="\n}`}
                    className="font-mono h-64 bg-amber-50 border-amber-200 rounded-xl"
                  />
                </CardContent>
              </Card>

              {/* Transaction Details Card */}
              <Card className="bg-white border border-amber-100 shadow-xl rounded-2xl overflow-hidden transition-all duration-300 hover:shadow-2xl">
                <CardHeader className="">
                  <CardTitle className="flex items-center gap-2 text-xl">
                    <Coffee className="h-5 w-5" />
                    Coffee Details
                  </CardTitle>
                  <CardDescription className="text-white/80">Support your favorite creator</CardDescription>
                </CardHeader>
                <CardContent className="pt-6">
                  {transaction ? (
                    <div className="space-y-4">
                      <div className="p-4 rounded-xl bg-gradient-to-r from-amber-50 to-yellow-50 border border-amber-200 flex items-center">
                        <div className="p-3 mr-4 bg-gradient-to-r from-amber-500 to-yellow-500 rounded-full">
                          <Coins className="h-5 w-5 text-white" />
                        </div>
                        <div>
                          <div className="text-sm text-amber-800/70 mb-1">Amount</div>
                          <div className="text-2xl font-bold text-amber-700">{transaction.amount}</div>
                        </div>
                      </div>

                      <div className="p-4 rounded-xl bg-gradient-to-r from-yellow-50 to-amber-50 border border-yellow-200">
                        <div className="text-sm text-amber-800/70 mb-1">Recipient (Creator)</div>
                        <div className="font-mono flex items-center justify-between">
                          <span className="text-amber-700">{transaction.recipient}</span>
                          <Button
                            variant="ghost"
                            size="icon"
                            onClick={() => copyToClipboard(transaction.recipient)}
                            className="text-amber-500 hover:text-amber-700 hover:bg-amber-50"
                          >
                            <Copy className="h-4 w-4" />
                          </Button>
                        </div>
                      </div>

                      <div className="p-4 rounded-xl bg-gradient-to-r from-amber-50 to-yellow-50 border border-amber-200">
                        <div className="text-sm text-amber-800/70 mb-1">Sender (You)</div>
                        <div className="font-mono flex items-center justify-between">
                          <span className="truncate text-amber-700">{formatAddress(transaction.sender)}</span>
                          <Button
                            variant="ghost"
                            size="icon"
                            onClick={() => copyToClipboard(transaction.sender)}
                            className="text-amber-500 hover:text-amber-700 hover:bg-amber-50"
                          >
                            <Copy className="h-4 w-4" />
                          </Button>
                        </div>
                      </div>

                      <div className="p-4 rounded-xl bg-gradient-to-r from-yellow-50 to-amber-50 border border-yellow-200">
                        <div className="text-sm text-amber-800/70 mb-1">Signature</div>
                        <div className="font-mono text-xs truncate text-amber-700">
                          {transaction.signature.substring(0, 20)}...
                        </div>
                      </div>
                    </div>
                  ) : (
                    <div className="h-64 flex flex-col items-center justify-center text-amber-300">
                      <div className="w-16 h-16 mb-4 rounded-full bg-amber-100 flex items-center justify-center">
                        <Coffee className="h-8 w-8 text-amber-300" />
                      </div>
                      <span className="text-amber-800/50">
                        {jsonInput ? "Invalid JSON format" : "Paste valid transaction JSON to buy a coffee"}
                      </span>
                    </div>
                  )}
                </CardContent>
                <CardFooter>
                  <Button
                    className="w-full py-6 text-lg rounded-xl bg-gradient-to-r from-amber-500 to-yellow-500 hover:from-amber-600 hover:to-yellow-600 shadow-lg hover:shadow-xl transition-all duration-300 hover:scale-[1.02]"
                    disabled={!isValid}
                    onClick={handleSubmit}
                  >
                    <Coffee className="mr-2 h-5 w-5" />
                    Buy Coffee
                  </Button>
                </CardFooter>
              </Card>
            </div>
          </TabsContent>

          <TabsContent value="history">
            <Card className="bg-white border border-amber-100 shadow-xl rounded-2xl overflow-hidden">
              <CardHeader className="">
                <CardTitle className="flex items-center gap-2 text-xl">
                  <Clock className="h-5 w-5" />
                  Coffee History
                </CardTitle>
                <CardDescription className="text-white/80">Your coffee support history</CardDescription>
              </CardHeader>
              <CardContent className="pt-6">
                {history.length > 0 ? (
                  <div className="space-y-4">
                    {history.map((tx, index) => (
                      <div
                        key={index}
                        className="p-4 rounded-xl bg-white border border-amber-100 shadow-md hover:shadow-lg transition-all duration-300 flex items-center"
                      >
                        <div className="mr-4 p-3 bg-gradient-to-r from-amber-400 to-yellow-500 rounded-full">
                          <Coffee className="h-5 w-5 text-white" />
                        </div>
                        <div className="flex-1">
                          <div className="flex justify-between items-start">
                            <div>
                              <div className="font-bold text-lg text-amber-800">{tx.amount}</div>
                              <div className="text-sm text-amber-700/70">To: {formatAddress(tx.recipient)}</div>
                            </div>
                            <div className="text-right">
                              <div className="text-amber-600 text-sm font-medium">Coffee Sent</div>
                              <div className="text-xs text-amber-700/50">
                                {tx.timestamp ? new Date(tx.timestamp).toLocaleString() : "Unknown date"}
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                ) : (
                  <div className="h-64 flex flex-col items-center justify-center text-amber-300">
                    <div className="w-16 h-16 mb-4 rounded-full bg-amber-100 flex items-center justify-center">
                      <Coffee className="h-8 w-8 text-amber-300" />
                    </div>
                    <span className="text-amber-800/50">No coffee purchases yet</span>
                  </div>
                )}
              </CardContent>
            </Card>
          </TabsContent>
        </Tabs>
      </div>
      <Toaster />
    </div>
  )
}
