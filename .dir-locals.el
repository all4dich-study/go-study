;; .dir-locals.el
((go-mode . ((eval . (progn
                       (setenv "GOOS" "darwin") ; Adjust as needed
                       (setenv "GOARCH" "arm64")   ; Adjust for your target MCU
                       (setenv "GOFLAGS" "-tags=tinygo")
                       ;; You might need to restart gopls or re-analyze for this to take effect
                       ;; Or set specific lsp client variables
                       )))))
