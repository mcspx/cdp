// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"github.com/mafredri/cdp/protocol/debugger"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/runtime"
)

// AddScriptToEvaluateOnLoadArgs represents the arguments for AddScriptToEvaluateOnLoad in the Page domain.
type AddScriptToEvaluateOnLoadArgs struct {
	ScriptSource string `json:"scriptSource"` // No description.
}

// NewAddScriptToEvaluateOnLoadArgs initializes AddScriptToEvaluateOnLoadArgs with the required arguments.
func NewAddScriptToEvaluateOnLoadArgs(scriptSource string) *AddScriptToEvaluateOnLoadArgs {
	args := new(AddScriptToEvaluateOnLoadArgs)
	args.ScriptSource = scriptSource
	return args
}

// AddScriptToEvaluateOnLoadReply represents the return values for AddScriptToEvaluateOnLoad in the Page domain.
type AddScriptToEvaluateOnLoadReply struct {
	Identifier ScriptIdentifier `json:"identifier"` // Identifier of the added script.
}

// AddScriptToEvaluateOnNewDocumentArgs represents the arguments for AddScriptToEvaluateOnNewDocument in the Page domain.
type AddScriptToEvaluateOnNewDocumentArgs struct {
	Source string `json:"source"` // No description.
}

// NewAddScriptToEvaluateOnNewDocumentArgs initializes AddScriptToEvaluateOnNewDocumentArgs with the required arguments.
func NewAddScriptToEvaluateOnNewDocumentArgs(source string) *AddScriptToEvaluateOnNewDocumentArgs {
	args := new(AddScriptToEvaluateOnNewDocumentArgs)
	args.Source = source
	return args
}

// AddScriptToEvaluateOnNewDocumentReply represents the return values for AddScriptToEvaluateOnNewDocument in the Page domain.
type AddScriptToEvaluateOnNewDocumentReply struct {
	Identifier ScriptIdentifier `json:"identifier"` // Identifier of the added script.
}

// CaptureScreenshotArgs represents the arguments for CaptureScreenshot in the Page domain.
type CaptureScreenshotArgs struct {
	// Format Image compression format (defaults to png).
	//
	// Values: "jpeg", "png".
	Format  *string   `json:"format,omitempty"`
	Quality *int      `json:"quality,omitempty"` // Compression quality from range [0..100] (jpeg only).
	Clip    *Viewport `json:"clip,omitempty"`    // Capture the screenshot of a given region only.
	// FromSurface Capture the screenshot from the surface, rather than
	// the view. Defaults to true.
	//
	// Note: This property is experimental.
	FromSurface *bool `json:"fromSurface,omitempty"`
}

// NewCaptureScreenshotArgs initializes CaptureScreenshotArgs with the required arguments.
func NewCaptureScreenshotArgs() *CaptureScreenshotArgs {
	args := new(CaptureScreenshotArgs)

	return args
}

// SetFormat sets the Format optional argument. Image compression
// format (defaults to png).
//
// Values: "jpeg", "png".
func (a *CaptureScreenshotArgs) SetFormat(format string) *CaptureScreenshotArgs {
	a.Format = &format
	return a
}

// SetQuality sets the Quality optional argument. Compression quality
// from range [0..100] (jpeg only).
func (a *CaptureScreenshotArgs) SetQuality(quality int) *CaptureScreenshotArgs {
	a.Quality = &quality
	return a
}

// SetClip sets the Clip optional argument. Capture the screenshot of
// a given region only.
func (a *CaptureScreenshotArgs) SetClip(clip Viewport) *CaptureScreenshotArgs {
	a.Clip = &clip
	return a
}

// SetFromSurface sets the FromSurface optional argument. Capture the
// screenshot from the surface, rather than the view. Defaults to true.
//
// Note: This property is experimental.
func (a *CaptureScreenshotArgs) SetFromSurface(fromSurface bool) *CaptureScreenshotArgs {
	a.FromSurface = &fromSurface
	return a
}

// CaptureScreenshotReply represents the return values for CaptureScreenshot in the Page domain.
type CaptureScreenshotReply struct {
	Data []byte `json:"data"` // Base64-encoded image data.
}

// CreateIsolatedWorldReply represents the return values for CreateIsolatedWorld in the Page domain.
type CreateIsolatedWorldReply struct {
	ExecutionContextID runtime.ExecutionContextID `json:"executionContextId"` // Execution context of the isolated world.
}

// GetAppManifestReply represents the return values for GetAppManifest in the Page domain.
type GetAppManifestReply struct {
	URL    string             `json:"url"`            // Manifest location.
	Errors []AppManifestError `json:"errors"`         // No description.
	Data   *string            `json:"data,omitempty"` // Manifest content.
}

// GetFrameTreeReply represents the return values for GetFrameTree in the Page domain.
type GetFrameTreeReply struct {
	FrameTree FrameTree `json:"frameTree"` // Present frame tree structure.
}

// GetLayoutMetricsReply represents the return values for GetLayoutMetrics in the Page domain.
type GetLayoutMetricsReply struct {
	LayoutViewport LayoutViewport `json:"layoutViewport"` // Metrics relating to the layout viewport.
	VisualViewport VisualViewport `json:"visualViewport"` // Metrics relating to the visual viewport.
	ContentSize    dom.Rect       `json:"contentSize"`    // Size of scrollable area.
}

// GetNavigationHistoryReply represents the return values for GetNavigationHistory in the Page domain.
type GetNavigationHistoryReply struct {
	CurrentIndex int               `json:"currentIndex"` // Index of the current navigation history entry.
	Entries      []NavigationEntry `json:"entries"`      // Array of navigation history entries.
}

// GetResourceContentReply represents the return values for GetResourceContent in the Page domain.
type GetResourceContentReply struct {
	Content       string `json:"content"`       // Resource content.
	Base64Encoded bool   `json:"base64Encoded"` // True, if content was served as base64.
}

// GetResourceTreeReply represents the return values for GetResourceTree in the Page domain.
type GetResourceTreeReply struct {
	FrameTree FrameResourceTree `json:"frameTree"` // Present frame / resource tree structure.
}

// HandleJavaScriptDialogArgs represents the arguments for HandleJavaScriptDialog in the Page domain.
type HandleJavaScriptDialogArgs struct {
	Accept     bool    `json:"accept"`               // Whether to accept or dismiss the dialog.
	PromptText *string `json:"promptText,omitempty"` // The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
}

// NewHandleJavaScriptDialogArgs initializes HandleJavaScriptDialogArgs with the required arguments.
func NewHandleJavaScriptDialogArgs(accept bool) *HandleJavaScriptDialogArgs {
	args := new(HandleJavaScriptDialogArgs)
	args.Accept = accept
	return args
}

// SetPromptText sets the PromptText optional argument. The text to
// enter into the dialog prompt before accepting. Used only if this is
// a prompt dialog.
func (a *HandleJavaScriptDialogArgs) SetPromptText(promptText string) *HandleJavaScriptDialogArgs {
	a.PromptText = &promptText
	return a
}

// NavigateToHistoryEntryArgs represents the arguments for NavigateToHistoryEntry in the Page domain.
type NavigateToHistoryEntryArgs struct {
	EntryID int `json:"entryId"` // Unique id of the entry to navigate to.
}

// NewNavigateToHistoryEntryArgs initializes NavigateToHistoryEntryArgs with the required arguments.
func NewNavigateToHistoryEntryArgs(entryID int) *NavigateToHistoryEntryArgs {
	args := new(NavigateToHistoryEntryArgs)
	args.EntryID = entryID
	return args
}

// PrintToPDFArgs represents the arguments for PrintToPDF in the Page domain.
type PrintToPDFArgs struct {
	Landscape               *bool    `json:"landscape,omitempty"`               // Paper orientation. Defaults to false.
	DisplayHeaderFooter     *bool    `json:"displayHeaderFooter,omitempty"`     // Display header and footer. Defaults to false.
	PrintBackground         *bool    `json:"printBackground,omitempty"`         // Print background graphics. Defaults to false.
	Scale                   *float64 `json:"scale,omitempty"`                   // Scale of the webpage rendering. Defaults to 1.
	PaperWidth              *float64 `json:"paperWidth,omitempty"`              // Paper width in inches. Defaults to 8.5 inches.
	PaperHeight             *float64 `json:"paperHeight,omitempty"`             // Paper height in inches. Defaults to 11 inches.
	MarginTop               *float64 `json:"marginTop,omitempty"`               // Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom            *float64 `json:"marginBottom,omitempty"`            // Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft              *float64 `json:"marginLeft,omitempty"`              // Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight             *float64 `json:"marginRight,omitempty"`             // Right margin in inches. Defaults to 1cm (~0.4 inches).
	PageRanges              *string  `json:"pageRanges,omitempty"`              // Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
	IgnoreInvalidPageRanges *bool    `json:"ignoreInvalidPageRanges,omitempty"` // Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
	HeaderTemplate          *string  `json:"headerTemplate,omitempty"`          // HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - `date`: formatted print date - `title`: document title - `url`: document location - `pageNumber`: current page number - `totalPages`: total pages in the document For example, `<span class=title></span>` would generate span containing the title.
	FooterTemplate          *string  `json:"footerTemplate,omitempty"`          // HTML template for the print footer. Should use the same format as the `headerTemplate`.
	PreferCSSPageSize       *bool    `json:"preferCSSPageSize,omitempty"`       // Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
}

// NewPrintToPDFArgs initializes PrintToPDFArgs with the required arguments.
func NewPrintToPDFArgs() *PrintToPDFArgs {
	args := new(PrintToPDFArgs)

	return args
}

// SetLandscape sets the Landscape optional argument. Paper
// orientation. Defaults to false.
func (a *PrintToPDFArgs) SetLandscape(landscape bool) *PrintToPDFArgs {
	a.Landscape = &landscape
	return a
}

// SetDisplayHeaderFooter sets the DisplayHeaderFooter optional argument.
// Display header and footer. Defaults to false.
func (a *PrintToPDFArgs) SetDisplayHeaderFooter(displayHeaderFooter bool) *PrintToPDFArgs {
	a.DisplayHeaderFooter = &displayHeaderFooter
	return a
}

// SetPrintBackground sets the PrintBackground optional argument.
// Print background graphics. Defaults to false.
func (a *PrintToPDFArgs) SetPrintBackground(printBackground bool) *PrintToPDFArgs {
	a.PrintBackground = &printBackground
	return a
}

// SetScale sets the Scale optional argument. Scale of the webpage
// rendering. Defaults to 1.
func (a *PrintToPDFArgs) SetScale(scale float64) *PrintToPDFArgs {
	a.Scale = &scale
	return a
}

// SetPaperWidth sets the PaperWidth optional argument. Paper width in
// inches. Defaults to 8.5 inches.
func (a *PrintToPDFArgs) SetPaperWidth(paperWidth float64) *PrintToPDFArgs {
	a.PaperWidth = &paperWidth
	return a
}

// SetPaperHeight sets the PaperHeight optional argument. Paper height
// in inches. Defaults to 11 inches.
func (a *PrintToPDFArgs) SetPaperHeight(paperHeight float64) *PrintToPDFArgs {
	a.PaperHeight = &paperHeight
	return a
}

// SetMarginTop sets the MarginTop optional argument. Top margin in
// inches. Defaults to 1cm (~0.4 inches).
func (a *PrintToPDFArgs) SetMarginTop(marginTop float64) *PrintToPDFArgs {
	a.MarginTop = &marginTop
	return a
}

// SetMarginBottom sets the MarginBottom optional argument. Bottom
// margin in inches. Defaults to 1cm (~0.4 inches).
func (a *PrintToPDFArgs) SetMarginBottom(marginBottom float64) *PrintToPDFArgs {
	a.MarginBottom = &marginBottom
	return a
}

// SetMarginLeft sets the MarginLeft optional argument. Left margin in
// inches. Defaults to 1cm (~0.4 inches).
func (a *PrintToPDFArgs) SetMarginLeft(marginLeft float64) *PrintToPDFArgs {
	a.MarginLeft = &marginLeft
	return a
}

// SetMarginRight sets the MarginRight optional argument. Right margin
// in inches. Defaults to 1cm (~0.4 inches).
func (a *PrintToPDFArgs) SetMarginRight(marginRight float64) *PrintToPDFArgs {
	a.MarginRight = &marginRight
	return a
}

// SetPageRanges sets the PageRanges optional argument. Paper ranges
// to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which
// means print all pages.
func (a *PrintToPDFArgs) SetPageRanges(pageRanges string) *PrintToPDFArgs {
	a.PageRanges = &pageRanges
	return a
}

// SetIgnoreInvalidPageRanges sets the IgnoreInvalidPageRanges optional argument.
// Whether to silently ignore invalid but successfully parsed page
// ranges, such as '3-2'. Defaults to false.
func (a *PrintToPDFArgs) SetIgnoreInvalidPageRanges(ignoreInvalidPageRanges bool) *PrintToPDFArgs {
	a.IgnoreInvalidPageRanges = &ignoreInvalidPageRanges
	return a
}

// SetHeaderTemplate sets the HeaderTemplate optional argument. HTML
// template for the print header. Should be valid HTML markup with
// following classes used to inject printing values into them: -
// `date`: formatted print date - `title`: document title - `url`:
// document location - `pageNumber`: current page number -
// `totalPages`: total pages in the document
//
// For example, `<span class=title></span>` would generate span
// containing the title.
func (a *PrintToPDFArgs) SetHeaderTemplate(headerTemplate string) *PrintToPDFArgs {
	a.HeaderTemplate = &headerTemplate
	return a
}

// SetFooterTemplate sets the FooterTemplate optional argument. HTML
// template for the print footer. Should use the same format as the
// `headerTemplate`.
func (a *PrintToPDFArgs) SetFooterTemplate(footerTemplate string) *PrintToPDFArgs {
	a.FooterTemplate = &footerTemplate
	return a
}

// SetPreferCSSPageSize sets the PreferCSSPageSize optional argument.
// Whether or not to prefer page size as defined by css. Defaults to
// false, in which case the content will be scaled to fit the paper
// size.
func (a *PrintToPDFArgs) SetPreferCSSPageSize(preferCSSPageSize bool) *PrintToPDFArgs {
	a.PreferCSSPageSize = &preferCSSPageSize
	return a
}

// PrintToPDFReply represents the return values for PrintToPDF in the Page domain.
type PrintToPDFReply struct {
	Data []byte `json:"data"` // Base64-encoded pdf data.
}

// ReloadArgs represents the arguments for Reload in the Page domain.
type ReloadArgs struct {
	IgnoreCache            *bool   `json:"ignoreCache,omitempty"`            // If true, browser cache is ignored (as if the user pressed Shift+refresh).
	ScriptToEvaluateOnLoad *string `json:"scriptToEvaluateOnLoad,omitempty"` // If set, the script will be injected into all frames of the inspected page after reload. Argument will be ignored if reloading dataURL origin.
}

// NewReloadArgs initializes ReloadArgs with the required arguments.
func NewReloadArgs() *ReloadArgs {
	args := new(ReloadArgs)

	return args
}

// SetIgnoreCache sets the IgnoreCache optional argument. If true,
// browser cache is ignored (as if the user pressed Shift+refresh).
func (a *ReloadArgs) SetIgnoreCache(ignoreCache bool) *ReloadArgs {
	a.IgnoreCache = &ignoreCache
	return a
}

// SetScriptToEvaluateOnLoad sets the ScriptToEvaluateOnLoad optional argument.
// If set, the script will be injected into all frames of the inspected
// page after reload. Argument will be ignored if reloading dataURL
// origin.
func (a *ReloadArgs) SetScriptToEvaluateOnLoad(scriptToEvaluateOnLoad string) *ReloadArgs {
	a.ScriptToEvaluateOnLoad = &scriptToEvaluateOnLoad
	return a
}

// RemoveScriptToEvaluateOnLoadArgs represents the arguments for RemoveScriptToEvaluateOnLoad in the Page domain.
type RemoveScriptToEvaluateOnLoadArgs struct {
	Identifier ScriptIdentifier `json:"identifier"` // No description.
}

// NewRemoveScriptToEvaluateOnLoadArgs initializes RemoveScriptToEvaluateOnLoadArgs with the required arguments.
func NewRemoveScriptToEvaluateOnLoadArgs(identifier ScriptIdentifier) *RemoveScriptToEvaluateOnLoadArgs {
	args := new(RemoveScriptToEvaluateOnLoadArgs)
	args.Identifier = identifier
	return args
}

// RemoveScriptToEvaluateOnNewDocumentArgs represents the arguments for RemoveScriptToEvaluateOnNewDocument in the Page domain.
type RemoveScriptToEvaluateOnNewDocumentArgs struct {
	Identifier ScriptIdentifier `json:"identifier"` // No description.
}

// NewRemoveScriptToEvaluateOnNewDocumentArgs initializes RemoveScriptToEvaluateOnNewDocumentArgs with the required arguments.
func NewRemoveScriptToEvaluateOnNewDocumentArgs(identifier ScriptIdentifier) *RemoveScriptToEvaluateOnNewDocumentArgs {
	args := new(RemoveScriptToEvaluateOnNewDocumentArgs)
	args.Identifier = identifier
	return args
}

// ScreencastFrameAckArgs represents the arguments for ScreencastFrameAck in the Page domain.
type ScreencastFrameAckArgs struct {
	SessionID int `json:"sessionId"` // Frame number.
}

// NewScreencastFrameAckArgs initializes ScreencastFrameAckArgs with the required arguments.
func NewScreencastFrameAckArgs(sessionID int) *ScreencastFrameAckArgs {
	args := new(ScreencastFrameAckArgs)
	args.SessionID = sessionID
	return args
}

// SearchInResourceReply represents the return values for SearchInResource in the Page domain.
type SearchInResourceReply struct {
	Result []debugger.SearchMatch `json:"result"` // List of search matches.
}

// SetAdBlockingEnabledArgs represents the arguments for SetAdBlockingEnabled in the Page domain.
type SetAdBlockingEnabledArgs struct {
	Enabled bool `json:"enabled"` // Whether to block ads.
}

// NewSetAdBlockingEnabledArgs initializes SetAdBlockingEnabledArgs with the required arguments.
func NewSetAdBlockingEnabledArgs(enabled bool) *SetAdBlockingEnabledArgs {
	args := new(SetAdBlockingEnabledArgs)
	args.Enabled = enabled
	return args
}

// SetBypassCSPArgs represents the arguments for SetBypassCSP in the Page domain.
type SetBypassCSPArgs struct {
	Enabled bool `json:"enabled"` // Whether to bypass page CSP.
}

// NewSetBypassCSPArgs initializes SetBypassCSPArgs with the required arguments.
func NewSetBypassCSPArgs(enabled bool) *SetBypassCSPArgs {
	args := new(SetBypassCSPArgs)
	args.Enabled = enabled
	return args
}

// SetFontFamiliesArgs represents the arguments for SetFontFamilies in the Page domain.
type SetFontFamiliesArgs struct {
	FontFamilies FontFamilies `json:"fontFamilies"` // Specifies font families to set. If a font family is not specified, it won't be changed.
}

// NewSetFontFamiliesArgs initializes SetFontFamiliesArgs with the required arguments.
func NewSetFontFamiliesArgs(fontFamilies FontFamilies) *SetFontFamiliesArgs {
	args := new(SetFontFamiliesArgs)
	args.FontFamilies = fontFamilies
	return args
}

// SetFontSizesArgs represents the arguments for SetFontSizes in the Page domain.
type SetFontSizesArgs struct {
	FontSizes FontSizes `json:"fontSizes"` // Specifies font sizes to set. If a font size is not specified, it won't be changed.
}

// NewSetFontSizesArgs initializes SetFontSizesArgs with the required arguments.
func NewSetFontSizesArgs(fontSizes FontSizes) *SetFontSizesArgs {
	args := new(SetFontSizesArgs)
	args.FontSizes = fontSizes
	return args
}

// SetDownloadBehaviorArgs represents the arguments for SetDownloadBehavior in the Page domain.
type SetDownloadBehaviorArgs struct {
	// Behavior Whether to allow all or deny all download requests, or use
	// default Chrome behavior if available (otherwise deny).
	//
	// Values: "deny", "allow", "default".
	Behavior     string  `json:"behavior"`
	DownloadPath *string `json:"downloadPath,omitempty"` // The default path to save downloaded files to. This is required if behavior is set to 'allow'
}

// NewSetDownloadBehaviorArgs initializes SetDownloadBehaviorArgs with the required arguments.
func NewSetDownloadBehaviorArgs(behavior string) *SetDownloadBehaviorArgs {
	args := new(SetDownloadBehaviorArgs)
	args.Behavior = behavior
	return args
}

// SetDownloadPath sets the DownloadPath optional argument. The
// default path to save downloaded files to. This is required if
// behavior is set to 'allow'
func (a *SetDownloadBehaviorArgs) SetDownloadPath(downloadPath string) *SetDownloadBehaviorArgs {
	a.DownloadPath = &downloadPath
	return a
}

// SetLifecycleEventsEnabledArgs represents the arguments for SetLifecycleEventsEnabled in the Page domain.
type SetLifecycleEventsEnabledArgs struct {
	Enabled bool `json:"enabled"` // If true, starts emitting lifecycle events.
}

// NewSetLifecycleEventsEnabledArgs initializes SetLifecycleEventsEnabledArgs with the required arguments.
func NewSetLifecycleEventsEnabledArgs(enabled bool) *SetLifecycleEventsEnabledArgs {
	args := new(SetLifecycleEventsEnabledArgs)
	args.Enabled = enabled
	return args
}

// StartScreencastArgs represents the arguments for StartScreencast in the Page domain.
type StartScreencastArgs struct {
	// Format Image compression format.
	//
	// Values: "jpeg", "png".
	Format        *string `json:"format,omitempty"`
	Quality       *int    `json:"quality,omitempty"`       // Compression quality from range [0..100].
	MaxWidth      *int    `json:"maxWidth,omitempty"`      // Maximum screenshot width.
	MaxHeight     *int    `json:"maxHeight,omitempty"`     // Maximum screenshot height.
	EveryNthFrame *int    `json:"everyNthFrame,omitempty"` // Send every n-th frame.
}

// NewStartScreencastArgs initializes StartScreencastArgs with the required arguments.
func NewStartScreencastArgs() *StartScreencastArgs {
	args := new(StartScreencastArgs)

	return args
}

// SetFormat sets the Format optional argument. Image compression
// format.
//
// Values: "jpeg", "png".
func (a *StartScreencastArgs) SetFormat(format string) *StartScreencastArgs {
	a.Format = &format
	return a
}

// SetQuality sets the Quality optional argument. Compression quality
// from range [0..100].
func (a *StartScreencastArgs) SetQuality(quality int) *StartScreencastArgs {
	a.Quality = &quality
	return a
}

// SetMaxWidth sets the MaxWidth optional argument. Maximum screenshot
// width.
func (a *StartScreencastArgs) SetMaxWidth(maxWidth int) *StartScreencastArgs {
	a.MaxWidth = &maxWidth
	return a
}

// SetMaxHeight sets the MaxHeight optional argument. Maximum
// screenshot height.
func (a *StartScreencastArgs) SetMaxHeight(maxHeight int) *StartScreencastArgs {
	a.MaxHeight = &maxHeight
	return a
}

// SetEveryNthFrame sets the EveryNthFrame optional argument. Send
// every n-th frame.
func (a *StartScreencastArgs) SetEveryNthFrame(everyNthFrame int) *StartScreencastArgs {
	a.EveryNthFrame = &everyNthFrame
	return a
}

// SetWebLifecycleStateArgs represents the arguments for SetWebLifecycleState in the Page domain.
type SetWebLifecycleStateArgs struct {
	// State Target lifecycle state
	//
	// Values: "frozen", "active".
	State string `json:"state"`
}

// NewSetWebLifecycleStateArgs initializes SetWebLifecycleStateArgs with the required arguments.
func NewSetWebLifecycleStateArgs(state string) *SetWebLifecycleStateArgs {
	args := new(SetWebLifecycleStateArgs)
	args.State = state
	return args
}

// SetProduceCompilationCacheArgs represents the arguments for SetProduceCompilationCache in the Page domain.
type SetProduceCompilationCacheArgs struct {
	Enabled bool `json:"enabled"` // No description.
}

// NewSetProduceCompilationCacheArgs initializes SetProduceCompilationCacheArgs with the required arguments.
func NewSetProduceCompilationCacheArgs(enabled bool) *SetProduceCompilationCacheArgs {
	args := new(SetProduceCompilationCacheArgs)
	args.Enabled = enabled
	return args
}

// AddCompilationCacheArgs represents the arguments for AddCompilationCache in the Page domain.
type AddCompilationCacheArgs struct {
	URL  string `json:"url"`  // No description.
	Data []byte `json:"data"` // Base64-encoded data
}

// NewAddCompilationCacheArgs initializes AddCompilationCacheArgs with the required arguments.
func NewAddCompilationCacheArgs(url string, data []byte) *AddCompilationCacheArgs {
	args := new(AddCompilationCacheArgs)
	args.URL = url
	args.Data = data
	return args
}

// GenerateTestReportArgs represents the arguments for GenerateTestReport in the Page domain.
type GenerateTestReportArgs struct {
	Message string  `json:"message"`         // Message to be displayed in the report.
	Group   *string `json:"group,omitempty"` // Specifies the endpoint group to deliver the report to.
}

// NewGenerateTestReportArgs initializes GenerateTestReportArgs with the required arguments.
func NewGenerateTestReportArgs(message string) *GenerateTestReportArgs {
	args := new(GenerateTestReportArgs)
	args.Message = message
	return args
}

// SetGroup sets the Group optional argument. Specifies the endpoint
// group to deliver the report to.
func (a *GenerateTestReportArgs) SetGroup(group string) *GenerateTestReportArgs {
	a.Group = &group
	return a
}
