package bedrock


// The type of harmful category usable in a content filter.
// Experimental.
type ContentFilterType string

const (
	// Describes input prompts and model responses that indicates sexual interest, activity, or arousal using direct or indirect references to body parts, physical traits, or sex.
	// Experimental.
	ContentFilterType_SEXUAL ContentFilterType = "SEXUAL"
	// Describes input prompts and model responses that includes glorification of or threats to inflict physical pain, hurt, or injury toward a person, group or thing.
	// Experimental.
	ContentFilterType_VIOLENCE ContentFilterType = "VIOLENCE"
	// Describes input prompts and model responses that discriminate, criticize, insult, denounce, or dehumanize a person or group on the basis of an identity (such as race, ethnicity, gender, religion, sexual orientation, ability, and national origin).
	// Experimental.
	ContentFilterType_HATE ContentFilterType = "HATE"
	// Describes input prompts and model responses that includes demeaning, humiliating, mocking, insulting, or belittling language.
	//
	// This type of language is also labeled
	// as bullying.
	// Experimental.
	ContentFilterType_INSULTS ContentFilterType = "INSULTS"
	// Describes input prompts and model responses that seeks or provides information about engaging in misconduct activity, or harming, defrauding, or taking advantage of a person, group or institution.
	// Experimental.
	ContentFilterType_MISCONDUCT ContentFilterType = "MISCONDUCT"
	// Enable to detect and block user inputs attempting to override system instructions.
	//
	// To avoid misclassifying system prompts as a prompt attack and ensure that the filters
	// are selectively applied to user inputs, use input tagging.
	// Experimental.
	ContentFilterType_PROMPT_ATTACK ContentFilterType = "PROMPT_ATTACK"
)

